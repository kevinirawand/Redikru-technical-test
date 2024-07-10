package web

type JobCreateRequest struct {
	CompanyId   string `json:"companyId" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
