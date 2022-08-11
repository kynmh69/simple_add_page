package util

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func makeHeader(req http.Request) {
	token := GetToken()
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	req.Header.Add("Notion-Version", "2022-06-28")
}

func PostRequest(url string, payload []byte) []byte {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	makeHeader(*req)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln("Request has error.")
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	return body
}
