package main

import (
	_ "encoding/json"
	_ "fmt"
	_ "os/user"
	_ "runtime/trace"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/lib/pq"
)

func main() {

	dsn := `host=localhost user=khurshid dbname=firstdb password=X sslmode=disable`

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	books := bookshop{
		Id: 11,
		Name:       "Alisher Navoiy",
		Capasity:   10000,
		
	}
	sect := aboutbookshop{1, 2, 2}

	// tx := db.Session(&gorm.Session{SkipDefaultTransaction:true})

	db.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&books).Error
		if err != nil {
			return err
		}
		err = tx.Create(&sect).Error
		if err != nil{
			return err
		}
		return nil
	})

}
