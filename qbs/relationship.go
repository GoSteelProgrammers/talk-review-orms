package main

import (
	"fmt"
	"github.com/coocood/qbs"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

//START STRUCT OMIT

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Age       int
	Profile   *Profile
	ProfileId int64
}

type Profile struct {
	Id        int64
	Homepage  string
	Interests string
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

	if err = migration.CreateTableIfNotExists(new(Profile)); err != nil {
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

	profile := &Profile{Homepage: "www.example.com", Interests: "Golang", Id: 1}
	q.Save(profile)
	q.Save(&User{Id: 1, FirstName: "John", LastName: "Doe", Age: 25, ProfileId: 1})

	//START CODE OMIT
	user = &User{Id: 1}
	q.Find(user)

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", user.Profile)
	//END CODE OMIT
}
