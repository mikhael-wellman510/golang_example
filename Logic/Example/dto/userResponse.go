package dto

import "time"

type UserResponse struct {
	Id        int
	UserName  string
	Age       int
	Address   string
	CreatedAt time.Time
}
