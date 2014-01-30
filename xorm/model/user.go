package model

type User struct {
	Id        int    `xorm:"not null pk autoincr integer"`
	FirstName string `xorm:"text"`
	LastName  string `xorm:"text"`
	Age       int    `xorm:"integer"`
}
