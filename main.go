package main

import (
	"Redikru-technical-test/app"
	"Redikru-technical-test/controller"
	"Redikru-technical-test/helper"
	"Redikru-technical-test/middleware"
	"Redikru-technical-test/repository"
	"Redikru-technical-test/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
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

	router := app.NewRouter(jobController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
