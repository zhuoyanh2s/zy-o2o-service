package models

import (
	"time"
)

type User struct {
	Id       int64
	Username string
	Password string
	State    int8
	IsStaff  bool
	IsAdmin  bool
	LastTime time.Time
	Created  time.Time
	Updated  time.Time
	//Profile  *Profile `orm:"null;rel(one)"`
}

type Profile struct {
	Id int64

	RealName string
	IdNumber string

	Age     int
	Email   string
	Mobile  string
	Address string
	Gender  string
}
