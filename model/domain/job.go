package domain

import "Redikru-technical-test/model/web"

type Job struct {
	Id          string                  `json:"id"`
	CompanyId   string                  `json:"companyId"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Company     web.CompanyBaseResponse `json:"company"`
}
