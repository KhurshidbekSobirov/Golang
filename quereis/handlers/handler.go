package handlers

import (
	"Golang/quereis/database"
	"Golang/quereis/models"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// var users = []models.User{
// 	{Id: 1, First_name: "Ali", Last_name: "Zairov"},
// 	{Id: 2, First_name: "Dili", Last_name: "Narzullayev"},
// 	{Id: 3, First_name: "Mohira", Last_name: "Doniyorov"},
// }

func Users(rw http.ResponseWriter, r *http.Request) {
	db := database.Con()
	rw.Header().Set("Content-Type", "Application/json")
	result, err := db.Query(`select * from users2`)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		var u = models.User{}
		err := result.Scan(&u.First_name)
		if err != nil {
			continue
		}
		fmt.Println(u)
		e := json.NewEncoder(rw)
     	e.Encode(u.First_name)
		 
}

// func Books(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "Application/json")
// 	books := []models.Book{
// 		{Name: "O'tkan kunlar", Price: 20000},
// 		{Name: "Dunyoni ishlari", Price: 21000},
// 		{Name: "Nano python", Price: 22000},
// 	}
// 	e := json.NewEncoder(rw)
// 	e.Encode(books)
// }

// func Get_Query(rw http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	e := json.NewEncoder(rw)
// 	for _, j := range users {
// 		id, err := strconv.Atoi(query["id"][0])
// 		if err != nil {
// 			panic(err)
// 		}
// 		if id == j.Id {
// 			e.Encode(j)
// 		}
// 	}
// }

// func QueryParams(w http.ResponseWriter, r *http.Request) {
// 	for k, v := range r.URL.Query() {
// 		fmt.Printf("%s: %s\n", k, v)
// 	}
// }
}