package test

import (
	"Redikru-technical-test/app"
	"Redikru-technical-test/controller"
	"Redikru-technical-test/helper"
	"Redikru-technical-test/middleware"
	"Redikru-technical-test/repository"
	"Redikru-technical-test/service"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func NewTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/TEST_redikru_technical_test")
	helper.PanicIfError(err)

	fmt.Println("TESTTTTT")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := NewTestDB()

	if db == nil {
		fmt.Println("ERRORR ANJENG")
	}
	validate := validator.New()
	JobRepository := repository.NewJobRepository()
	jobService := service.NewJobService(JobRepository, db, validate)
	jobController := controller.NewJobController(jobService)

	router := app.NewRouter(jobController)

	return middleware.NewAuthMiddleware(router)
}

func TestGetJobsSuccess(t *testing.T) {
	router := setupRouter()

	params := url.Values{}
	params.Add("title", "VIN")
	params.Add("description", "test")
	params.Add("companyName", "KAN")

	url := fmt.Sprintf("http://localhost:3000/api/jobs?%s", params.Encode())

	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}
func TestCreateJobSuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"companyId": "b974744b-3ecc-11ef-a580-04d4c4de4bcb",
  "title": "KEVIN", "description": "from test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/jobs", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}
func TestUnauthorized(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"companyId": "007b43a6-3e80-11ef-a580-04d4c4de4bcb",
  "title": "KEVIN", "description": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/jobs", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "BUKAN RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestCreateJobFailed(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"companyId": "007b43a6-3e80-11ef-a580-04d4c4de4bcb",
  "title": "KEVIN", "description": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/jobs", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}
