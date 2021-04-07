package model

type User struct {
	UserId   string `json:user_id`
	Name     string `json:name`
	Email    string `json:email`
	Password string `json:password`
}
