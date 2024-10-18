package entities

type Employers struct {
	ID                string `json:"id"`
	Image             string `json:"image"`
	OrganizationName  string `json:"organization_name"`
	INN               string `json:"inn"`
	Email             string `json:"email"`
	AboutOrganization string `json:"bio"`
	Date              string `json:"date"`
	Phone             string `json:"phone"`
	Password          string `json:"password"`
	Token             string `json:"token"`
	Created_at        string `json:"created_At"`
	Updated_at        string `json:"updated_At"`
}
