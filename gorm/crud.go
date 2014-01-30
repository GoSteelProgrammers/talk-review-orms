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

func PrintTable(db gorm.DB) {
	var users []User

	db.Find(&users)

	fmt.Printf("%+v\n", users)
}

func main() {
	var (
		db   gorm.DB
		user *User
	)

	user = new(User)

	db = SetupDb()

	//START CODE OMIT
	db.Save(&User{FirstName: "John", LastName: "Doe", Age: 25})
	PrintTable(db)

	db.First(user, 1)

	user.FirstName = "James"
	db.Save(user)
	PrintTable(db)

	db.Delete(user)
	PrintTable(db)
	//END CODE OMIT
}
