package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	ConfigFilename string = "../config.json"
)

type Config struct {
	ServicePort int    `json:"serviceport"`
	RedisDSN    string `json:"redisdsn"`
	MysqlDSN    string `json:"mysqldsn"`
	ENV         string `json:"env"`
	RequestLog  string `json:"requestlog"`
	ErrorLog    string `json:"errorlog"`
}

func LoadConfig(filename *string) Config {

	var config Config

	configFile, err := os.Open(*filename)
	defer configFile.Close()

	if err != nil {
		fmt.Println(gin.DefaultWriter, err)
	}

	json.NewDecoder(configFile).Decode(&config)
	return config

}

func ConfigLoader() Config {

	var config Config

	_, err := os.Stat(ConfigFilename)
	if os.IsExist(err) {
		log.Println(err)
	}
	configFile, _ := os.Open(ConfigFilename)
	defer configFile.Close()

	if err != nil {
		fmt.Println(gin.DefaultWriter, err)
	}

	json.NewDecoder(configFile).Decode(&config)
	return config

}
