package _interface

import (
	"Redikru-technical-test/model/web"
	"golang.org/x/net/context"
)

type ICompanyService interface {
	FindAll(ctx context.Context) []web.CompanyBaseResponse
}
