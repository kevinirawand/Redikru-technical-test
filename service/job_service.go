package service

import (
	"Redikru-technical-test/helper"
	"Redikru-technical-test/model/domain"
	"Redikru-technical-test/model/web"
	_repositoryInterface "Redikru-technical-test/repository/interface"
	_serviceInterface "Redikru-technical-test/service/interface"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"net/http"
)

type JobService struct {
	JobRepository _repositoryInterface.IJobRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewJobService(jobRepository _repositoryInterface.IJobRepository, DB *sql.DB, validate *validator.Validate) _serviceInterface.IJobService {
	return &JobService{
		JobRepository: jobRepository,
		DB:            DB,
		Validate:      validate,
	}
}

func (service *JobService) Create(ctx context.Context, request web.JobCreateRequest) web.JobBaseResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	job := domain.Job{
		CompanyId:   request.CompanyId,
		Title:       request.Title,
		Description: request.Description,
	}

	job = service.JobRepository.Save(ctx, tx, job)

	return helper.ToJobResponse(job)
}

func (service *JobService) FindAll(ctx context.Context, request *http.Request) []web.JobBaseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	jobs := service.JobRepository.FindAll(ctx, tx, request)

	var jobsResponse []web.JobBaseResponse

	for _, job := range jobs {
		helper.PanicIfError(err)
		jobsResponse = append(jobsResponse, helper.ToJobResponse(job))
	}

	fmt.Println(jobs)

	return jobsResponse
}
