package repository

import (
	"Redikru-technical-test/helper"
	"Redikru-technical-test/model/domain"
	"Redikru-technical-test/model/web"
	_interface "Redikru-technical-test/repository/interface"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"net/http"
)

type JobRepository struct {
}

func NewJobRepository() _interface.IJobRepository {
	return &JobRepository{}
}

func (repository *JobRepository) Save(ctx context.Context, tx *sql.Tx, job domain.Job) domain.Job {
	newUUID := uuid.New()
	job.Id = newUUID.String()
	binaryUUID, err := newUUID.MarshalBinary()

	SQL := "INSERT INTO jobs (id, company_id, title, description) VALUES (?, ?, ?, ?)"

	_, err = tx.ExecContext(ctx, SQL, binaryUUID, job.CompanyId, job.Title, job.Description)

	helper.PanicIfError(err)

	return job
}

func (repository *JobRepository) FindAll(ctx context.Context, tx *sql.Tx, request *http.Request) []domain.Job {
	var queryParams []interface{}

	sqlSelect := "SELECT BIN_TO_UUID(A.Id) as Id, A.company_id as company_id, A.title as title, A.description as description, BIN_TO_UUID(B.id) as M_company_id, B.name as company_name "
	sqlFrom := "FROM jobs AS A JOIN companies AS B ON BIN_TO_UUID(B.Id) = A.company_id "
	sqlCondition := "WHERE 1=1 "

	if title := request.URL.Query().Get("title"); title != "" {
		sqlCondition += " AND A.title LIKE ?"
		queryParams = append(queryParams, "%"+title+"%")
	}
	if description := request.URL.Query().Get("description"); description != "" {
		sqlCondition += " AND A.description LIKE ?"
		queryParams = append(queryParams, "%"+description+"%")
	}
	if companyName := request.URL.Query().Get("companyName"); companyName != "" {
		sqlCondition += " AND B.name LIKE ?"
		queryParams = append(queryParams, "%"+companyName+"%")
	}

	sqlQuery := sqlSelect + sqlFrom + sqlCondition
	fmt.Println(sqlQuery)

	rows, err := tx.QueryContext(ctx, sqlQuery, queryParams...)

	helper.PanicIfError(err)

	var jobs []domain.Job

	for rows.Next() {
		job := domain.Job{}
		var companyId string
		var companyName string
		err := rows.Scan(&job.Id, &job.CompanyId, &job.Title, &job.Description, &companyId, &companyName)

		helper.PanicIfError(err)

		job.Company = web.CompanyBaseResponse{
			Id:   companyId,
			Name: companyName,
		}

		jobs = append(jobs, job)
	}

	return jobs
}
