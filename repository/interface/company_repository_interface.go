package _interface

import (
	"Redikru-technical-test/model/domain"
	"context"
	"database/sql"
)

type ICompanyRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Company
}
