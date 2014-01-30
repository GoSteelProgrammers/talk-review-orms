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
    last_name VARCHAR(80)
  );
  `); err != nil {
		panic(err)
	}

	if col, err = sess.Collection("users"); err != nil {
		panic(err)
	}

	return col
}

func PrintTable(col db.Collection) {
	var users []User

	col.Find().All(&users)

	fmt.Printf("%+v\n", users)
}

func main() {
	var (
		col  db.Collection
		user User
	)

	col = SetupCollection()

	//START CODE OMIT
	col.Append(User{Id: 1, FirstName: "John", LastName: "Doe"})
	PrintTable(col)

	col.Find(db.Cond{"id": 1}).One(&user)

	user.FirstName = "James"
	col.Find(user.Id).Update(user)
	PrintTable(col)

	col.Find(db.Cond{"id": user.Id}).Remove()
	PrintTable(col)
	//END CODE OMIT
}
