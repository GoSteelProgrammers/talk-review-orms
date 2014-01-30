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

func PrintTable(dbMap *gorp.DbMap) {
	var users []User

	dbMap.Select(&users, "SELECT * FROM users")

	fmt.Printf("%+v\n", users)
}

func main() {
	var (
		dbMap *gorp.DbMap
		user  *User
		iUser interface{}
	)

	dbMap = SetupDb()

	//START CODE OMIT
	dbMap.Insert(&User{Id: 1, FirstName: "John", LastName: "Doe"})
	PrintTable(dbMap)

	iUser, _ = dbMap.Get(User{}, 1)
	user = iUser.(*User)

	user.FirstName = "James"
	dbMap.Update(user)
	PrintTable(dbMap)

	dbMap.Delete(user)
	PrintTable(dbMap)
	//END CODE OMIT
}
