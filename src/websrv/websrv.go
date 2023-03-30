package websrv

import (
	"fmt"
	"net/http"
	"pizzeria-management-service/src/config"
	"pizzeria-management-service/src/tracer"
	"time"
)

var WebServer app_Server

type app_Server struct {
	startTime int64
}

func (srv *app_Server) Init() {
	go srv.runListener()
}

func (srv *app_Server) runListener() {
	defer func() {
		if err := recover(); err != nil {
			tracer.Error("websrv.runListener", "Critical error termination", fmt.Sprintf("%v", err))
		}
	}()

	tracer.Debug("Starting web server")

	port := ":" + config.Settings.WebServer.Port

	serverMux := http.NewServeMux()
	HttpServer := http.Server{
		Addr:        port,
		Handler:     serverMux,
		IdleTimeout: 20 * time.Second,
	}
	serverMux.HandleFunc("/", srv.route)
	serverMux.HandleFunc("/access/", srv.routeAccess)
	if e := HttpServer.ListenAndServe(); e != nil && e != http.ErrServerClosed {
		tracer.Error("websrv.runListener", "Critical error termination", e.Error())
	}
}
