package entities

type Reward struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Type       string `json:"type"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	About      string `json:"about"`
	Created_at string `json:"created_At"`
}
