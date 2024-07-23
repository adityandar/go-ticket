package request

type UpdateOrganizerRequestBody struct {
	Id          int
	FullName    string `json:"full_name"`
	Email       string
	CompanyName string `json:"company_name"`
}
