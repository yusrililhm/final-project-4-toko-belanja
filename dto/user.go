package dto

import "time"

type CreateNewUsersRequest struct {
	FullName string `json:"full_name" valid:"required~Full name can't be empty" example:"Harry Maguire"`
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"maguire.harry@mufc.com"`
	Password string `json:"password" valid:"required~Password can't be empty, length(6|255)~Minimum password is 6 length" example:"secret"`
}

type UsersLoginRequest struct {
	Email    string `json:"email" valid:"required~Email can't be empty, email" example:"maguire.harry@mufc.com"`
	Password string `json:"password" valid:"required~Password can't be empty, length(6|255)~Minimum password is 6 length" example:"secret"`
}

type CreateNewUsersResponse struct {
	Id        int       `json:"id" example:"1"`
	FullName  string    `json:"full_name" example:"Harry Maguire"`
	Email     string    `json:"email" example:"maguire.harry@mufc.com"`
	Password  string    `json:"password" example:"hash password"`
	Balance   uint      `json:"balance" example:"120000"`
	CreatedAt time.Time `json:"created_at" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserUpdateResponse struct {
	Id        int       `json:"id" example:"1"`
	FullName  string    `json:"full_name" example:"Harry Maguire"`
	Email     string    `json:"email" example:"maguire.harry@mufc.com"`
	Password  string    `json:"password" example:"hash password"`
	Balance   uint      `json:"balance" example:"120000"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-10-09T05:14:35.19324086+07:00"`
}

type UsersTopUpRequest struct {
	Balance uint `json:"balance" valid:"required~ Balance can't be empty, range(0|100000000)~ Balance can't be less than 0 or more than 100,000,000" example:"150000"`
}

type UsersLoginResponse struct {
	Token   string `json:"token"`
}

type UserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type AdminResponse struct {
	Message string `json:"message"`
}