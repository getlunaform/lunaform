package restapi

import (
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/backend/database"
	operation "github.com/getlunaform/lunaform/restapi/operations/providers"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func DeleteTfProviderConfigurationController(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operation.DeleteProviderConfigurationHandlerFunc {
	return func(params operation.DeleteProviderConfigurationParams, user *models.ResourceAuthUser) middleware.Responder {
		ch.SetRequest(params.HTTPRequest)

		errCode, err := buildDeleteTfProviderConfigurationResponse(db, params.ID)
		if err != nil {
			return NewServerErrorResponse(errCode, err.Error())
		}
		return operation.NewDeleteProviderConfigurationNoContent()
	}
}

func buildDeleteTfProviderConfigurationResponse(db database.Database, configId string) (errCode int, err error) {
	if err := db.Delete(DB_TABLE_TF_PROVIDER_CONFIGURATION, configId); err != nil {
		if _, notFound := err.(database.RecordDoesNotExistError); notFound {
			errCode = http.StatusNotFound
		} else {
			errCode = http.StatusInternalServerError
		}
	}
	return
}
