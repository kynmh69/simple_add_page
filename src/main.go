package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://api.notion.com/v1/pages"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer secret_8huHwdrEnHB6MNzU5BJ1peF4KyqvFhUPKUAfvg8205p")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
