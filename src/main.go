package main

import (
	"log"
	"simple_add_page/src/notion/page"
	"simple_add_page/src/util"
)

func main() {
	logger := log.Default()

	content := util.ReadPhotoIdFile("")
	id_list := util.SplitLines(*content)

	tag_name := "hiroki"

	for _, v := range id_list {
		page.Create(v, tag_name)
		logger.Println("Created. id:", v)
	}

}
