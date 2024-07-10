package web

type JobBaseResponse struct {
	Id          string              `json:"id"`
	CompanyId   string              `json:"companyId"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Company     CompanyBaseResponse `json:"company"`
}
