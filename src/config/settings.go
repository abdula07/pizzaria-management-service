package config

import (
	"encoding/json"
	"os"
	"pizzeria-management-service/src/tracer"
)

var Settings app_config

type app_config struct {
	WebServer struct {
		Port  string `json:"port"`
		Token string `json:"token"`
	} `json:"webserver"`
	DbMain struct {
		Login        string `json:"login"`
		Password     string `json:"password"`
		DatabaseName string `json:"databaseName"`
	}
}

func (obj *app_config) Init() bool {
	tracer.Debug("Init settings")
	if obj.load() {
		tracer.Debug("Settings inited")
		return true
	} else {
		tracer.Error("Fail to init settings")
		return false
	}
}

func (obj *app_config) load() bool {
	fileName, e := os.Getwd()
	if e != nil {
		tracer.Error(e.Error())
		return false
	}
	fileName += "/settings.json"
	if _, e := os.Stat(fileName); e != nil {
		tracer.Error(e.Error())
		return false
	}
	config, e := os.ReadFile(fileName)
	if e != nil {
		tracer.Error(e.Error())
		return false
	}
	e = json.Unmarshal(config, &obj)
	if e != nil {
		tracer.Error(e.Error())
		return false
	}
	return true
}
