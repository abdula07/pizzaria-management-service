package websrv

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pizzeria-management-service/src/config"
	"pizzeria-management-service/src/services"
	"pizzeria-management-service/src/tracer"
	"strings"
)

func (srv *app_Server) route(w http.ResponseWriter, r *http.Request) {

	data := WebResponse{}
	rUri := strings.Split(r.RequestURI, "/")
	if len(rUri) < 3 {
		srv.writeResponse(w, http.StatusBadRequest, data)
		return
	}
	switch rUri[1] {
	case "orders":
		switch rUri[2] {
		case "create":
			body, e := ioutil.ReadAll(r.Body)
			if e != nil {
				data.Message = e.Error()
			} else {
				var products []services.OrderProducts
				e = json.Unmarshal(body, &products)
				if e != nil {
					data.Message = e.Error()
				} else {
					data = WebResponse{
						Success: true,
						Message: "",
						Data:    services.Orders.Orders(products),
					}
				}
			}
		case "items":
			body, e := ioutil.ReadAll(r.Body)
			if e != nil {
				data.Message = e.Error()
			} else {
				var products []services.OrderProducts
				e = json.Unmarshal(body, &products)
				if e != nil {
					data.Message = e.Error()
				} else {
					data = WebResponse{
						Success: true,
						Message: "",
						Data:    services.Orders.Products(products, rUri[3]),
					}
				}
			}
		case "get":
			data = WebResponse{
				Success: true,
				Message: "",
				Data:    services.Orders.Get(rUri[3]),
			}

		default:
			data.Message = "Incorrect url"
		}
	default:
		data.Message = "Incorrect url"
	}

	srv.writeResponse(w, http.StatusOK, data)
}

func (srv *app_Server) routeAccess(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	tracer.Debug("Call >>>", token, ">>>", r.Method, r.RequestURI)

	data := WebResponse{}
	rUri := strings.Split(r.RequestURI, "/")
	if token != config.Settings.WebServer.Token {
		srv.writeResponse(w, http.StatusUnauthorized, data)
		return
	}
	switch rUri[2] {
	case "orders":
		switch rUri[3] {
		case "done":
			data = WebResponse{
				Success: true,
				Message: "",
				Data:    services.Orders.DoneMake(rUri[4]),
			}
		case "all":
			param := rUri[4]
			var done bool
			isCondition := false
			if param == "1" {
				done = true
				isCondition = true
			}
			if param == "0" {
				done = false
				isCondition = true
			}
			data = WebResponse{
				Success: true,
				Message: "",
				Data:    services.Orders.All(done, isCondition),
			}
		default:
			data.Message = "Incorrect url"
		}
	default:
		data.Message = "Incorrect url"
	}
	srv.writeResponse(w, http.StatusOK, data)
}

func (srv *app_Server) writeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	dataJs, err := json.Marshal(data)
	if err != nil {
		tracer.Error("router.writeResponse", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(dataJs)
	if err != nil {
		tracer.Error("router.writeResponse", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
