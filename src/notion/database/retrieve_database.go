package database

import (
	"simple_add_page/src/util"
)

func Retrieve() []byte {

	url := "https://api.notion.com/v1/databases/database_id"

	body := util.GetRequest(url)

	return body
}
