package entity

type UserEntity struct {
	UserId     int64  `json:"user_id"`
	UserName   string `json:"user_name" validate:"required"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	UserStatus string `json:"user_status"`
	Department string `json:"department" validate:"required"`
}
