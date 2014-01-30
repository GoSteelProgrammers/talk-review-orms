package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

//START STRUCT OMIT

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Age       int
}

//END STRUCT OMIT

func SetupDb() (db gorm.DB) {
	var (
		err error
	)

	if db, err = gorm.Open("sqlite3", ":memory:"); err != nil {
		panic(err)
	}

	db.CreateTable(User{})

	return db
}

func main() {
	var (
		db gorm.DB
	)

	db = SetupDb()

	db.Save(&User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       24,
	})

	db.Save(&User{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       52,
	})
	db.Save(&User{
		FirstName: "Joe",
		LastName:  "Shmoe",
		Age:       10,
	})

	//START CODE OMIT
	var users []User
	db.Where("last_name = ?", "Doe").Or("age < ?", 12).Limit(2).Order("age", true).Find(&users)
	for _, user := range users {
		fmt.Printf("%+v,", user)
	}
	fmt.Println()
	//END CODE OMIT
}
