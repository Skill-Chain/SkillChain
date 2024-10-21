package entities

import "time"

type User struct {
	ID         string    `json:"id"`
	Image      string    `json:"image"`
	Firstname  string    `json:"firstname"`
	Lastname   string    `json:"Surname"`
	Email      string    `json:"email"`
	Bio        string    `json:"bio"`
	Date       string    `json:"date"`
	Phone      string    `json:"phone"`
	Rewards    []Reward  `json:"rewards"`
	Skills     []Skill   `json:"skills"`
	Token      string    `json:"token"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_At"`
	Updated_at time.Time `json:"updated_At"`
}
