package restapi

import (
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/getlunaform/lunaform/backend/identity"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"

	"github.com/getlunaform/lunaform/helpers"
	operations "github.com/getlunaform/lunaform/restapi/operations/modules"
	"net/http"
)

func CreateTfModuleController(idp identity.Provider, ch *helpers.ContextHelper, db database.Database) operations.CreateModuleHandlerFunc {
	return operations.CreateModuleHandlerFunc(func(params operations.CreateModuleParams, p *models.ResourceAuthUser) (r middleware.Responder) {
		ch.SetRequest(params.HTTPRequest)

		buildTfModuleControllerCreateResponse(params.TerraformModule, db, ch)

		return operations.NewCreateModuleCreated().WithPayload(params.TerraformModule)
	})
}

func buildTfModuleControllerCreateResponse(module *models.ResourceTfModule, db database.Database, ch *helpers.ContextHelper) (errCode int, err error) {

	module.ID = idGenerator.MustGenerate()
	if err := db.Create(DB_TABLE_TF_MODULE, module.ID, module); err != nil {
		return http.StatusInternalServerError, err
	}
	module.Links = helpers.HalRootRscLinks(ch)
	module.Embedded = nil

	return
}
