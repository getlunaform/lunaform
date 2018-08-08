package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"fmt"
)

func DeleteTfProviderController(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operation.DeleteProviderHandlerFunc {
	return func(params operation.DeleteProviderParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		if errCode, err := buildDeleteTfProviderResponse(db, params.Name); err != nil {
			return NewServerErrorResponse(errCode, err.Error())
		}

		return operation.NewDeleteProviderNoContent()
	}
}

func buildDeleteTfProviderResponse(db database.Database, providerName string) (errCode int, err error) {
	if err = db.Delete(DB_TABLE_TF_PROVIDER, providerName); err != nil {
		if _, notFound := err.(database.RecordDoesNotExistError); notFound {
			errCode = http.StatusNotFound
			err = fmt.Errorf("could not find provider with name '%s'", providerName)
		} else {
			errCode = http.StatusInternalServerError
		}
	}
	return
}
