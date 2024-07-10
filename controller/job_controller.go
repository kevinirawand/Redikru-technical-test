package controller

import (
	"Redikru-technical-test/helper"
	"Redikru-technical-test/model/web"
	_serviceInterface "Redikru-technical-test/service/interface"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type JobController struct {
	JobService _serviceInterface.IJobService
}

func NewJobController(jobService _serviceInterface.IJobService) *JobController {
	return &JobController{JobService: jobService}
}

func (controller *JobController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jobCreateRequest := web.JobCreateRequest{}
	helper.ReadFromRequestBody(request, &jobCreateRequest)

	jobResponse := controller.JobService.Create(request.Context(), jobCreateRequest)
	response := web.Response{
		Code:   200,
		Status: "OK",
		Data:   jobResponse,
	}

	helper.WriteToResponseBody(writer, response)
}

func (controller *JobController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jobResponse := controller.JobService.FindAll(request.Context(), request)
	response := web.Response{
		Code:   200,
		Status: "OK",
		Data:   jobResponse,
	}

	helper.WriteToResponseBody(writer, response)
}
