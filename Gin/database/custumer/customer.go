package custumer

import (
	"app/database"
	"app/models"
)

func Create(costumer models.Costumer) error {

	db := database.Conn()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Commit()
	_, err = tx.Exec(`INSERT INTO
	 customers(name,password,money) 
	 VALUES($1,$2,$3)`, costumer.Name, costumer.Password, costumer.Money)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	return err
}

func Login(costumer models.Costumer) (models.Costumer2, error) {
	var new_costumer models.Costumer2

	db := database.Conn()
	defer db.Close()

	res := db.QueryRow(`SELECT name,money FROM customers WHERE name=$1 AND password=$2`, costumer.Name, costumer.Password)
	err := res.Scan(
		&new_costumer.Name,
		&new_costumer.Money,
	)
	if err != nil {
		panic(err)
	}
	new_costumer.Money = costumer.Money
	rows, err := db.Query(`SELECT id,name,price FROM products`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int64
		var str string
		var price int64
		err = rows.Scan(
			&id,
			&str,
			&price,
		)
		if err != nil {
			panic(err)
		}
		for _, j := range costumer.Products {
			if j == int(id) {
				new_costumer.Products = append(new_costumer.Products, str)
				new_costumer.Money -= price
				break
			}
		}
	}
	return new_costumer, nil
}

func Query_id(costumer models.Costumer) (models.Costumer2, error) {
	var new_costumer models.Costumer2

	db := database.Conn()
	defer db.Close()

	res := db.QueryRow(`SELECT id,name,money FROM customers WHERE id=$1`, costumer.Id)
	err := res.Scan(
		&new_costumer.Id,
		&new_costumer.Name,
		&new_costumer.Money,
	)
	if err != nil {
		panic(err)
	}

	var prd []string

	for _, j := range costumer.Products {
		var (
			str   string
			price int64
		)
		res := db.QueryRow(`SELECT name,price FROM products WHERE id=$1`, j)
		err = res.Scan(
			&str,
			&price,
		)
		if err != nil {
			panic(err)
		}
		prd = append(prd, str)
		new_costumer.Money -= price
	}

	_, err = db.Exec(`INSERT INTO customers(products) VALUES($1)`, prd)
	if err != nil {
		panic(err)
	}

	return new_costumer, nil
}
