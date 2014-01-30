package main

import (
	"database/sql"
	"fmt"
	"upper.io/db"
	_ "upper.io/db/sqlite"
)

//START STRUCT OMIT

type User struct {
	Id        int    `field:"id"`
	FirstName string `field:"first_name"`
	LastName  string `field:"last_name"`
	Age       int    `field:"age"`
}

//END STRUCT OMIT

func SetupCollection() (col db.Collection) {
	var (
		err      error
		sess     db.Database
		settings db.Settings
	)

	settings = db.Settings{
		Database: `:memory:`,
	}

	if sess, err = db.Open("sqlite", settings); err != nil {
		panic(err)
	}

	if _, err = sess.Driver().(*sql.DB).Exec(`CREATE TABLE users (
    id INTEGER,
    first_name VARCHAR(80),
    last_name VARCHAR(80),
    age INTEGER
  );
  `); err != nil {
		panic(err)
	}

	if col, err = sess.Collection("users"); err != nil {
		panic(err)
	}

	return col
}

func main() {
	var (
		col db.Collection
	)

	col = SetupCollection()

	//START SETUP OMIT
	col.Append(User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Age:       24,
	})

	col.Append(User{
		Id:        2,
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       52,
	})

	col.Append(User{
		Id:        3,
		FirstName: "Joe",
		LastName:  "Shmoe",
		Age:       10,
	})
	//END SETUP OMIT

	//START CODE OMIT

	var users []User
	col.Find(db.Or{
		db.Cond{"last_name": "Doe"},
		db.Cond{"age <": "12"},
	}).
		Limit(2).
		Sort("-age").
		All(&users)

	fmt.Printf("%+v\n", users)

	//END CODE OMIT
}
