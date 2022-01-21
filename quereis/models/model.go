package models

type User struct {
 Id int `json:"id"`
 First_name string `json:"first_name"`
 Last_name  string `json:"last_name"`
}
type Book struct {
 Name  string `json:"name"`
 Price int64  `json:"price"`
}