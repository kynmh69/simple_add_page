package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"simple_add_page/src/model"
)

func main() {

	url := "https://api.notion.com/v1/pages"

	create_page := model.New("test", "hiroki")
	json_payload, err := create_page.Marshal()

	if err != nil {
		log.Fatalln("json error.")
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(json_payload))

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer secret_8huHwdrEnHB6MNzU5BJ1peF4KyqvFhUPKUAfvg8205p")
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln("Request has error.")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
