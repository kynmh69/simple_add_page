package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"simple_add_page/src/model"
	"simple_add_page/src/util"
)

func main() {

	logger := log.Default()

	url := "https://api.notion.com/v1/pages"

	create_page := model.New("test", "hiroki")
	json_payload, err := create_page.Marshal()

	if err != nil {
		log.Fatalln("json error.")
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(json_payload))

	token := getToken()

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln("Request has error.")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	logger.Println(res)
	logger.Println(string(body))

}

func getToken() string {
	conf := util.ReadConfigFile("")
	token := "Bearer "
	token += conf.ApiSettings.ApiSecret
	return token
}
