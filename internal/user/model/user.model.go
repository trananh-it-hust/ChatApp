package model

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRegister struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

func (u *UserRegister) TableName() string {
	return (&User{}).TableName()
}

type UserLogin struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u *UserLogin) TableName() string {
	return (&User{}).TableName()
}

type UserLoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (u *UserLoginResponse) TableName() string {
	return (&User{}).TableName()
}
