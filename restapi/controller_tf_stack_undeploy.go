package restapi

import (
	"fmt"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/backend/workers"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/models"
	operations "github.com/getlunaform/lunaform/restapi/operations/stacks"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func DeleteTfStackController(
	idp identity.Provider, ch *helpers.ContextHelper,
	db database.Database,
	workerPool *workers.TfAgentPool,
) operations.UndeployStackHandlerFunc {
	return operations.UndeployStackHandlerFunc(func(params operations.UndeployStackParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		stack := &models.ResourceTfStack{}
		if err := db.Read(DB_TABLE_TF_STACK, params.ID, stack); err != nil {
			if _, stackNotFound := err.(database.RecordDoesNotExistError); stackNotFound {
				return operations.NewUndeployStackNoContent()
			} else {
				return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
			}
		}

		// Make sure we remove the stack from the parent module
		module := models.ResourceTfModule{}
		if err := db.Read(DB_TABLE_TF_MODULE, module.ID, &module); err != nil {
			if _, stackNotFound := err.(database.RecordDoesNotExistError); stackNotFound {
				return NewServerErrorResponse(http.StatusBadRequest, fmt.Sprintf(
					"Could not find module '%s' for stack '%s'", module.ID, stack.ID))
			}
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		}

		for i, moduleStack := range module.Embedded.Stacks {
			if moduleStack.ID == stack.ID {
				module.Embedded.Stacks = append(module.Embedded.Stacks[:i], module.Embedded.Stacks[i+1:]...)
				break
			}
		}
		if err := db.Update(DB_TABLE_TF_MODULE, module.ID, module); err != nil {
			return NewServerErrorResponse(http.StatusInternalServerError, err.Error())
		}

		db.Delete(DB_TABLE_TF_STACK, stack.ID)

		return NewServerErrorResponse(http.StatusInternalServerError, "Could not delete stack.")
	})
}
