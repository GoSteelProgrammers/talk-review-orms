package main

import (
	"fmt"
	"github.com/lunny/xorm"
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

func SetupDb() (engine *xorm.Engine) {
	var (
		err error
	)

	if engine, err = xorm.NewEngine("sqlite3", ":memory:"); err != nil {
		panic(err)
	}

	if err = engine.Sync(new(User)); err != nil {
		panic(err)
	}

	return engine
}

func PrintTable(engine *xorm.Engine) {
	var users []User

	engine.Find(&users)

	fmt.Printf("%+v\n", users)
}

func main() {
	var (
		engine *xorm.Engine
		user   *User
	)

	engine = SetupDb()

	//START CODE OMIT
	engine.Insert(&User{Id: 1, FirstName: "John", LastName: "Doe"})
	PrintTable(engine)

	user = &User{Id: 1}
	engine.Get(user)

	user.FirstName = "James"
	engine.Update(user)
	PrintTable(engine)

	engine.Delete(user)
	PrintTable(engine)
	//END CODE OMIT
}
