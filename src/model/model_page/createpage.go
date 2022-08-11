// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//	createPage, err := UnmarshalCreatePage(bytes)
//	bytes, err = createPage.Marshal()
package modelpage

import (
	"encoding/json"
	"simple_add_page/src/util"
)

func UnmarshalCreatePage(data []byte) (CreatePage, error) {
	var r CreatePage
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CreatePage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func New(content_title string, multi_select_name string) (c *CreatePage) {
	conf := util.ReadConfigFile("")

	multi_select := MultiSelect{Name: multi_select_name}
	multi_select_arr := [...]MultiSelect{multi_select}
	tag := Tag{ID: conf.ApiSettings.TagId, Type: util.TYPE_MULTI_SELECT, MultiSelect: multi_select_arr[:]}

	text := Text{Content: content_title}

	title := Title{Type: "text", Text: text}
	title_arr := [...]Title{title}
	photo_id := PhotoID{ID: "title", Type: "title", Title: title_arr[:]}

	propaties := Properties{PhotoID: photo_id, Tag: tag}
	parent := Parent{DatabaseID: conf.ApiSettings.DatabaseID}

	return &CreatePage{Parent: parent, Properties: propaties}

}

type CreatePage struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	PhotoID PhotoID `json:"photo_id"`
	Tag     Tag     `json:"tag"`
}

type PhotoID struct {
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

type Tag struct {
	ID          string        `json:"id"`
	Type        string        `json:"type"`
	MultiSelect []MultiSelect `json:"multi_select"`
}

type MultiSelect struct {
	Name string `json:"name"`
}
