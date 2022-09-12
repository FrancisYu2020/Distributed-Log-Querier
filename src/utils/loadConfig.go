package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Server struct {
	IpAddr   string `json:ipaddr`
	Name     string `json:name`
	Port     string `json:port`
	FilePath string `json:filePath`
}

func getConfigs() string {

	b, err := ioutil.ReadFile("/home/hangy6/mp1-hangy6-tian23/config.json") // pass the file name

	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
	return str
}

func deserializeJson(configJson string) []Server {

	jsonAsBytes := []byte(configJson)
	configs := make([]Server, 0)
	err := json.Unmarshal(jsonAsBytes, &configs)
	if err != nil {
		panic(err)
	}

	return configs
}

func LoadConfig() []Server {

	// Unmarshal config component into a slice of structs.
	jsonConfigList := getConfigs()
	unmarshelledConfigs := deserializeJson(jsonConfigList)
	return unmarshelledConfigs
}

// reference: https://blog.csdn.net/pzython/article/details/113716598
