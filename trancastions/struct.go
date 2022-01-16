package main


type bookshop struct{
	Id int64 `json:"id"`
	Name string `json:"name"`
	Capasity int64 `json:"capasity"`
	// Section_name []string `json:"section_name"`
	// Outlet_id []int64 `json:"outlet_id"`
	// Outlet_name []string `json:"outlet_name"`
}

// type section struct{
// 	id int64 `json:"id"`
// 	Name string `json:"name"`
//}

type aboutbookshop struct{
	shop_id int64
	section_id int64
	outline_id int64
}