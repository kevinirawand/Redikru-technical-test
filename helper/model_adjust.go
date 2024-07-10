package helper

import (
	"Redikru-technical-test/model/domain"
	"Redikru-technical-test/model/web"
)

func ToJobResponse(job domain.Job) web.JobBaseResponse {
	return web.JobBaseResponse{
		Id:          job.Id,
		CompanyId:   job.CompanyId,
		Title:       job.Title,
		Description: job.Description,
		Company:     job.Company,
	}
}
