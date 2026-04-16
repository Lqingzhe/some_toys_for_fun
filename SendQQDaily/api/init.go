package api

import (
	"harassment/model"
	"os"

	"gopkg.in/yaml.v3"
)

func GetInfo() model.Config {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	newStruct := model.Config{}
	err = yaml.Unmarshal(data, &newStruct)
	if err != nil {
		panic(err)
	}
	return newStruct
}
