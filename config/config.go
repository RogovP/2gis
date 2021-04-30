package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database DB `json:"db"`
	Server Server `json:"server"`
}

type Server struct {
	Port string `json:"port"`
}

type DB struct {
	Host string `json:"host"`
	Name string `json:"name"`
	User string `json:"user"`
	Password string `json:"password"`
	Port string `json:"port"`
}

func GetConfig() Config {
	var config Config
	var configFile, err = os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	var jsonParser = json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}