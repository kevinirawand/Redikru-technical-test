package repository

import (
	"Redikru-technical-test/helper"
	"Redikru-technical-test/model/domain"
	_interface "Redikru-technical-test/repository/interface"
	"context"
	"database/sql"
)

type CompanyRepository struct {
}

func NewCompanyRepository() _interface.ICompanyRepository {
	return &CompanyRepository{}
}

func (repository *CompanyRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Company {
	SQL := "SELECT * FROM companies"

	rows, err := tx.QueryContext(ctx, SQL)

	helper.PanicIfError(err)

	var companies []domain.Company

	for rows.Next() {
		company := domain.Company{}
		err := rows.Scan(&company.Id, &company.Name)
		helper.PanicIfError(err)
		companies = append(companies, company)
	}

	return companies
}
