package page

import (
	"log"
	"simple_add_page/src/model"
	"simple_add_page/src/util"
)

func Create(page_name string, tag string) {
	logger := log.Default()
	url := "https://api.notion.com/v1/pages"

	create_page := model.New(page_name, tag)
	json_payload, err := create_page.Marshal()

	if err != nil {
		log.Fatalln("json error.")
	}

	body := util.PostRequest(url, json_payload)

	logger.Println(string(body))
}
