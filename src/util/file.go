package util

func ReadConfigFile(config_file_path string) {
	if config_file_path != nil || conconfig_file_path != "" {
		
	} else {
		config_file_path = "./config/api.yml"
	}
	
}

type ApiSettings struct {
	ApiSecret string
	DatabaseID string

}