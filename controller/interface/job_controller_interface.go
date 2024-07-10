package _interface

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IJobController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
