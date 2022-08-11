// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    database, err := UnmarshalDatabase(bytes)
//    bytes, err = database.Marshal()

package modeldatabase

import "encoding/json"

func UnmarshalDatabase(data []byte) (Database, error) {
	var r Database
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Database) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Database struct {
	Object         string         `json:"object"`
	ID             string         `json:"id"`
	Cover          interface{}    `json:"cover"`
	Icon           interface{}    `json:"icon"`
	CreatedTime    string         `json:"created_time"`
	CreatedBy      TedBy          `json:"created_by"`
	LastEditedBy   TedBy          `json:"last_edited_by"`
	LastEditedTime string         `json:"last_edited_time"`
	Title          []TitleElement `json:"title"`
	Description    []interface{}  `json:"description"`
	IsInline       bool           `json:"is_inline"`
	Properties     Properties     `json:"properties"`
	Parent         Parent         `json:"parent"`
	URL            string         `json:"url"`
	Archived       bool           `json:"archived"`
}

type TedBy struct {
	Object string `json:"object"`
	ID     string `json:"id"`
}

type Parent struct {
	Type   string `json:"type"`
	PageID string `json:"page_id"`
}

type Properties struct {
	Tag     Tag     `json:"tag"`
	PhotoID PhotoID `json:"photo_id"`
}

type PhotoID struct {
	ID    string       `json:"id"`
	Name  string       `json:"name"`
	Type  string       `json:"type"`
	Title PhotoIDTitle `json:"title"`
}

type PhotoIDTitle struct {
}

type Tag struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	MultiSelect MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Options []Option `json:"options"`
}

type Option struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type TitleElement struct {
	Type        string      `json:"type"`
	Text        Text        `json:"text"`
	Annotations Annotations `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        interface{} `json:"href"`
}

type Annotations struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type Text struct {
	Content string      `json:"content"`
	Link    interface{} `json:"link"`
}
