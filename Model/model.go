package model

import "time"

type Student struct {
	ID          int       `json:"id"`
	Firstname   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         int       `json:"age"`
	Email       string    `json:"email"`
	Marks       int       `json:"marks"`
	JoiningDate time.Time `json:"joining_date"`
	Result      string    `json:"result"`
}
