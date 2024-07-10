package _interface

import (
	"Redikru-technical-test/model/web"
	"golang.org/x/net/context"
	"net/http"
)

type IJobService interface {
	Create(ctx context.Context, request web.JobCreateRequest) web.JobBaseResponse
	FindAll(ctx context.Context, request *http.Request) []web.JobBaseResponse
}
