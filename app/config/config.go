package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	DB            string `json:"db"`
	Redis         string `json:"redis"`
	Kafka         string `json:"kafka"`
	Elasticsearch struct {
		Addresses []string `json:"addresses"`
		IndexName string   `json:"index_name"`
	} `json:"elasticsearch"`
}

func ReadConfigAndArg() *Config {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fileConfig := "config.json"
	// log.Println("basepath: ", basepath+"/"+fileConfig)
	data, err := os.ReadFile(basepath + "/" + fileConfig)
	if err != nil {
		log.Fatalln(err)
	}
	// var tempCfg *Config
	var tempCfg Config
	if data != nil {
		err = json.Unmarshal(data, &tempCfg)
		if err != nil {
			log.Fatalf("Unmarshal err %v", err.Error())
		}
	}
	return &tempCfg
}
