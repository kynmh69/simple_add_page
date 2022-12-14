package util

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func ReadConfigFile(config_file_path string) *Config {
	// logger := log.Default()

	if config_file_path == "" {
		config_file_path = "../config/api.yml"
	}

	// dir, _ := ioutil.ReadDir("../")
	// for i := 0; i < len(dir); i++ {
	// 	logger.Println("file path: ", dir[i].Name())
	// }

	f, err := os.Open(config_file_path)

	if err != nil {
		log.Fatalln("read config file error.", err)
	}
	defer f.Close()

	var conf Config
	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		log.Fatalln("yaml decode error", err)
	}
	return &conf
}

func ReadLines(file_path string) []string {
	var lines []string
	if file_path == "" {
		file_path = "../input_file/photo_list.txt"
	}
	content, _ := ioutil.ReadFile(file_path)
	log.Default().Println(string(content))
	f, err := os.Open(file_path)
	if err != nil {
		log.Fatalln("can't read photo id list.", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ReplaceAll(text, ".jpg", "")
		log.Default().Println(text)
		lines = append(lines, text)
	}
	return lines
}

func GetPrjDir() string {
	return "../"
}

type ApiSettings struct {
	ApiSecret  string `yaml:"api_secret"`
	DatabaseID string `yaml:"database_id"`
	TagId      string `yaml:"tag_id"`
}

type Config struct {
	ApiSettings ApiSettings `yaml:"api_settings"`
}
