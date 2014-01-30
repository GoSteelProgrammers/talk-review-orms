package main

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

//START STRUCT OMIT

type User struct {
	Id        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int    `db:"age"`
}

//END STRUCT OMIT
func SetupDb() (dbMap *gorp.DbMap) {
	var (
		err error
		db  *sql.DB
	)

	if db, err = sql.Open("sqlite3", ":memory:"); err != nil {
		panic(nil)
	}

	dbMap = &gorp.DbMap{
		Db:      db,
		Dialect: gorp.SqliteDialect{},
	}

	//START REGISTER OMIT

	dbMap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	//END REGISTER OMIT

	if err = dbMap.CreateTablesIfNotExists(); err != nil {
		panic(err)
	}

	return dbMap
}

func main() {
	var (
		dbMap *gorp.DbMap
	)

	dbMap = SetupDb()

	//START SETUP OMIT
	dbMap.Insert(&User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Age:       24,
	})

	dbMap.Insert(&User{
		Id:        2,
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       52,
	})

	dbMap.Insert(&User{
		Id:        3,
		FirstName: "Joe",
		LastName:  "Shmoe",
		Age:       10,
	})
	//END SETUP OMIT

	//START CODE OMIT

	var users []User
	dbMap.Select(&users, "SELECT * FROM users WHERE last_name = ? OR age < ? ORDER BY age DESC LIMIT 2", "Doe", 12)
	fmt.Printf("%+v\n", users)

	//END CODE OMIT
}
