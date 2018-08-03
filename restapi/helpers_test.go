package restapi

import "github.com/getlunaform/lunaform/backend/database"

func getTestingDB(content []map[string]string) database.Database {
	dbDriver, err := database.NewMemoryDBDriverWithCollection(content)
	if err != nil {
		panic(err)
	}
	
	return database.NewDatabaseWithDriver(dbDriver)
}
