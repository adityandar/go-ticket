package structs

type Organizer struct {
	UserId      int    `json:"user_id"`
	CompanyName string `json:"company_name"`
	User        User
}
