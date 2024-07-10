package _interface

import (
	"Redikru-technical-test/model/domain"
	"database/sql"
	"golang.org/x/net/context"
	"net/http"
)

type IJobRepository interface {
	Save(ctx context.Context, tx *sql.Tx, job domain.Job) domain.Job
	FindAll(ctx context.Context, tx *sql.Tx, request *http.Request) []domain.Job
}
