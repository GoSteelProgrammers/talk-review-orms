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

	if engine, err = xorm.NewEngine("sqlite3", "/tmp/tmp.db"); err != nil {
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
	)

	engine = SetupDb()

	engine.Insert(&User{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Age:       24,
	})

	engine.Insert(&User{
		Id:        2,
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       52,
	})

	engine.Insert(&User{
		Id:        3,
		FirstName: "Joe",
		LastName:  "Shmoe",
		Age:       10,
	})

	//START CODE OMIT

	var users []User
	engine.Where("last_name = ? OR age < ?", "Doe", "12").OrderBy("age").Limit(2).Find(&users)
	fmt.Printf("%+v\n", users)

	//END CODE OMIT
}
