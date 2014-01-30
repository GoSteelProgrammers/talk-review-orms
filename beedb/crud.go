package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/mattn/go-sqlite3"
)

//START STRUCT OMIT

type User struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
}

//END STRUCT OMIT

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

func PrintTable(orm beedb.Model) {
	var users []User

	if err := orm.FindAll(&users); err != nil {
		panic(err)
	}

	fmt.Printf("%+v,", users)
	for _, user := range users {
		fmt.Printf("%+v,", user)
	}
	fmt.Println()
}

func main() {
	var (
		orm  beedb.Model
		user *User
	)

	user = new(User)

	orm = SetupDb()

	//START CODE OMIT
	orm.SetTable("users").Insert(map[string]interface{}{"id": 1, "first_name": "John", "last_name": "Doe", "age": 25})
	PrintTable(orm)

	orm.Where(1).Find(user)

	user.FirstName = "James"
	orm.Save(user)
	PrintTable(orm)

	orm.Delete(user)
	PrintTable(orm)
	//END CODE OMIT
}
