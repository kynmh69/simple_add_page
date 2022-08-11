package main

import (
	"log"
	"simple_add_page/src/notion/database"
	"simple_add_page/src/notion/page"
	"simple_add_page/src/util"
)

func main() {
	retrieveDatabase()
	createPage()
	retrieveDatabase()
}

func createPage() {
	logger := log.Default()

	content := util.ReadPhotoIdFile("")
	id_list := util.SplitLines(*content)

	tag_name := "hiroki"

	for _, v := range id_list {
		page.Create(v, tag_name)
		logger.Println("Created. id:", v)
	}
}

func retrieveDatabase() {
	logger := log.Default()

	body := database.Retrieve()

	logger.Println("database:", body)
	logger.Println("title:", body.Title[0].Text.Content)
}
