package app

import (
	_interface "Redikru-technical-test/controller/interface"
	"Redikru-technical-test/execption"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(jobController _interface.IJobController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/jobs", jobController.FindAll)
	router.POST("/api/jobs", jobController.Create)
	router.PanicHandler = execption.ErrorHandler

	return router
}
