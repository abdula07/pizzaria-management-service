package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"pizzeria-management-service/src/config"
	dbMain "pizzeria-management-service/src/dbMain"
	"pizzeria-management-service/src/tracer"
	"pizzeria-management-service/src/websrv"
	"syscall"
)

func main() {
	defer func() {
		if ce := recover(); ce != nil {
			log.Println("Critical error", fmt.Sprintf("%s", ce))
			os.Exit(1)
		}
	}()

	tracer.Debug("Staring app")

	if !config.Settings.Init() {
		os.Exit(1)
	}

	if !dbMain.ConnectToDb() {
		os.Exit(1)
	}

	websrv.WebServer.Init()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	tracer.Debug("Closing app")

}
