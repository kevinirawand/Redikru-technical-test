package main

import (
	"Redikru-technical-test/app"
	"Redikru-technical-test/controller"
	"Redikru-technical-test/execption"
	"Redikru-technical-test/helper"
	"Redikru-technical-test/middleware"
	"Redikru-technical-test/repository"
	"Redikru-technical-test/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := app.NewDB()

	if db == nil {
		fmt.Println("ERRORR ANJENG")
	}
	validate := validator.New()
	JobRepository := repository.NewJobRepository()
	jobService := service.NewJobService(JobRepository, db, validate)
	jobController := controller.NewJobController(jobService)
	router := httprouter.New()

	router.GET("/api/jobs", jobController.FindAll)
	router.POST("/api/jobs", jobController.Create)

	router.PanicHandler = execption.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
