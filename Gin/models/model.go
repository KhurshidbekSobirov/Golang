package models

type Costumer struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Money int64 `json:"money"`
	Products []int `json:"product"`
}

type Costumer2 struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Money int64 `json:"money"`
	Products []string `json:"products"`
}



type Product struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Price int64 `json:"price"`
}