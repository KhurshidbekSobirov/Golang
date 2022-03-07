package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegisterResponseModel struct {
	Message string `json:"message"`
}

type VerifyResponseModel struct {
	ID           string `json:"id"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserData struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Code      string `json:"code"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ListUsers struct {
	Users []User `json:"users"`
}

type GetProfileByJwtRequestModel struct{
	Token string `json:"token"`
}

// Validate Set Password Model
// func (spn *User) Validate() error {
// 	return validation.ValidateStruct(
// 		spn,
// 		validation.Field(&spn.NewPassword, validation.Required, validation.Length(8, 30), Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
// 	)
// }

// Validate Register User Model
func (rum *User) Validate() error {
	return validation.ValidateStruct(
		rum,
		validation.Field(&rum.FirstName, validation.Required, validation.Length(1, 30)),
		validation.Field(&rum.LastName, validation.Required, validation.Length(1, 30)),
		validation.Field(&rum.Username, validation.Required, validation.Length(5, 30), validation.Match(regexp.MustCompile("^[0-9a-z_.]+$"))),
		validation.Field(&rum.Email, validation.Required, is.Email),
		validation.Field(&rum.Password, validation.Required, validation.Length(8, 30), Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
	)
}


