package main

import (
	"fmt"
	"github.com/coocood/qbs"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

//START STRUCT OMIT

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Age       int
}

//END STRUCT OMIT

func init() {
	os.Remove("/tmp/tmp.db")
	qbs.Register("sqlite3", "/tmp/tmp.db", "", qbs.NewSqlite3())
}

func SetupDb() (q *qbs.Qbs) {
	var (
		err       error
		migration *qbs.Migration
	)

	qbs.SetLogger(log.New(os.Stdout, "query:", log.Lmicroseconds), log.New(os.Stdout, "error:", log.Lmicroseconds))

	q, err = qbs.GetQbs()

	if q, err = qbs.GetQbs(); err != nil {
		panic(err)
	}

	if migration, err = qbs.GetMigration(); err != nil {
		panic(err)
	}
	defer migration.Close()
	if err = migration.CreateTableIfNotExists(new(User)); err != nil {
		panic(err)
	}

	return q
}

func PrintTable(q *qbs.Qbs) {
	var users []*User

	q.FindAll(&users)

	for _, user := range users {
		fmt.Printf("%+v,", user)
	}
	fmt.Println()
}

func main() {
	var (
		q    *qbs.Qbs
		user *User
	)

	q = SetupDb()
	defer q.Close()

	//START CODE OMIT
	q.Save(&User{Id: 1, FirstName: "John", LastName: "Doe", Age: 25})
	PrintTable(q)

	user = &User{Id: 1}
	q.Find(user)

	user.FirstName = "James"
	q.Save(user)
	PrintTable(q)

	q.Delete(user)
	PrintTable(q)
	//END CODE OMIT
}
