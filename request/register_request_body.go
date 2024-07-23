package request

type RegisterRequestBody struct {
	FullName    string `json:"full_name"`
	Email       string
	Password    string
	CompanyName string `json:"company_name"`
}
