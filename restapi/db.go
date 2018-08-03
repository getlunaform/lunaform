package restapi

import "github.com/getlunaform/lunaform/backend/database"

const (
	DB_TABLE_TF_WORKSPACE    = database.DBTableRecordType("lf-workspace")
	DB_TABLE_TF_MODULE       = database.DBTableRecordType("lf-module")
	DB_TABLE_TF_STACK        = database.DBTableRecordType("lf-stack")
	DB_TABLE_TF_STATEBACKEND = database.DBTableRecordType("lf-statebackend")
	DB_TABLE_AUTH_USER       = database.DBTableRecordType("lf-auth-user")
)