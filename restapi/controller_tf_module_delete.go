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

		module := &models.ResourceTfModule{}
		if err := db.Read(DB_TABLE_TF_MODULE, params.ID, module); err != nil {
			if _, moduleNotFound := err.(database.RecordDoesNotExistError); moduleNotFound {
				return operations.NewDeleteModuleNoContent()
			} else {
				return NewServerError(http.StatusInternalServerError, err.Error())
			}
		}

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

		db.Delete(DB_TABLE_TF_MODULE, params.ID)

		return NewServerError(http.StatusInternalServerError, "Could not delete module")
	})
}
