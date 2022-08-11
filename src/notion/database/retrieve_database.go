package database

import (
	"log"
	modeldatabase "simple_add_page/src/model/model_database"
	"simple_add_page/src/util"
)

func Retrieve() *modeldatabase.Database {
	config := util.ReadConfigFile("")
	url := "https://api.notion.com/v1/databases/" + config.ApiSettings.DatabaseID

	body := util.GetRequest(url)
	database_struct, err := modeldatabase.UnmarshalDatabase(body)

	if err != nil {
		log.Fatalln("Can't unmarshal database.")
	}

	return &database_struct
}
