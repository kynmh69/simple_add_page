// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    createPage, err := UnmarshalCreatePage(bytes)
//    bytes, err = createPage.Marshal()

package model

import "encoding/json"

func UnmarshalCreatePage(data []byte) (CreatePage, error) {
	var r CreatePage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CreatePage struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	写真の番号 写真の番号 `json:"写真の番号"`
	タグ    タグ    `json:"タグ"`
}

type タグ struct {
	ID          string        `json:"id"`
	Type        string        `json:"type"`
	MultiSelect []MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Name string `json:"name"`
}

type 写真の番号 struct {
	ID    string  `json:"id"`
	Type  string  `json:"type"`
	Title []Title `json:"title"`
}

type Title struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}
