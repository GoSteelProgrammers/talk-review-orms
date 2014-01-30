package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
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
//
func init() {
	beedb.PluralizeTableNames = true
}

func SetupDb() (orm beedb.Model) {
	var (
		err error
		db  *sql.DB
	)

	if db, err = sql.Open("sqlite3", ":memory:"); err != nil {
		panic(err)
	}

	if _, err = db.Exec(`CREATE TABLE users (
    id INTEGER,
    first_name VARCHAR(80),
    last_name VARCHAR(80),
    age INTEGER
  );
  `); err != nil {
		panic(err)
	}

	return beedb.New(db)
}

func main() {
	var (
		orm beedb.Model
	)

	orm = SetupDb()

	//START SETUP OMIT
	orm.SetTable("users").Insert(map[string]interface{}{
		"id":         1,
		"first_name": "John",
		"last_name":  "Doe",
		"age":        24,
	})
	orm.SetTable("users").Insert(map[string]interface{}{
		"id":         2,
		"first_name": "Jane",
		"last_name":  "Doe",
		"age":        52,
	})
	orm.SetTable("users").Insert(map[string]interface{}{
		"id":         1,
		"first_name": "Joe",
		"last_name":  "Shmoe",
		"age":        10,
	})
	//END SETUP OMIT

	//START CODE OMIT

	var users []User
	orm.Where("last_name = ? OR age < ?", "Doe", 12).Limit(2).OrderBy("age").FindAll(&users)
	for _, user := range users {
		fmt.Printf("%+v,", user)
	}
	fmt.Println()

	//END CODE OMIT
}
