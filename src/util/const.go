package util

const TYPE_MULTI_SELECT = "multi_select"

func GetToken() string {
	conf := ReadConfigFile("")
	token := "Bearer "
	token += conf.ApiSettings.ApiSecret
	return token
}
