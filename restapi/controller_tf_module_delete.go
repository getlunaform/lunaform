package restapi

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/go-openapi/runtime/middleware"

	"github.com/getlunaform/lunaform/helpers"
	operations "github.com/getlunaform/lunaform/restapi/operations/modules"
	"strings"
	"fmt"
	"net/http"
)

var DeleteTfModuleController = func(idp identity.Provider, ch helpers.ContextHelper, db database.Database) operations.DeleteModuleHandlerFunc {
	return operations.DeleteModuleHandlerFunc(func(params operations.DeleteModuleParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		if err := buildDeleteTfModuleResponse(db, params.ID); err == nil {
			r = operations.NewDeleteModuleNoContent()
		}

		return

	})
}

var buildDeleteTfModuleResponse = func(db database.Database, moduleId string) (*CommonServerErrorResponder) {
	module := &models.ResourceTfModule{}
	if err := db.Read(DB_TABLE_TF_MODULE, moduleId, module); err != nil {
		if _, moduleNotFound := err.(database.RecordDoesNotExistError); moduleNotFound {
			return nil
		} else {
			return NewServerError(http.StatusInternalServerError, err.Error())
		}
	}

	if module.Embedded != nil {
		if len(module.Embedded.Stacks) > 0 {
			stackIds := make([]string, 0)
			for _, stk := range module.Embedded.Stacks {
				stackIds = append(stackIds, stk.ID)
			}
			return NewServerError(
				http.StatusUnprocessableEntity,
				fmt.Sprintf(
					"Could not delete module as it is relied up by stacks ['%s']",
					strings.Join(stackIds, "','"),
				),
			)
		}
	}

	if err := db.Delete(DB_TABLE_TF_MODULE, moduleId); err != nil {
		return NewServerError(http.StatusUnprocessableEntity, err.Error())
	}

	return nil
}
