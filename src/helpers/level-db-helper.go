package helpers

import (
	"bytes"
	"encoding/json"
	"pizzeria-management-service/src/tracer"
)

func ObjectToByte(obj any) []byte {
	objAsBytes := new(bytes.Buffer)
	err := json.NewEncoder(objAsBytes).Encode(obj)
	if err != nil {
		tracer.Error(err.Error())
	}
	return objAsBytes.Bytes()
}
