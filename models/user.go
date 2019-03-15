package models

import (
	"errors"
	"log"
	"time"
	"zy-o2o-service/utils"
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
	//Profile  *Profile `engine:"null;rel(one)"`
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

func AddUser(u *User) (bool, error) {
	pw := utils.Md5(u.Password)
	u.Password = pw

	_, err := engine.Insert(u)
	if err != nil {
		log.Println("保存用户数据到数据库时失败 =>", err)
		return false, errors.New("保存用户失败")
	}
	return true, nil
}

func DeleteUser(id int64) (bool, error) {
	user := User{Id: id}
	_, err := engine.Delete(&user)
	if err != nil {
		log.Println("删除用户数据保存到数据库时失败 =>", err)
		return false, errors.New("删除用户失败")
	}
	return true, nil

}

func UpdateUser(u *User) (user *User, err error) {
	_, err = engine.Update(u)
	if err != nil {
		log.Println("更新用户数据保存到数据库时失败 =>", err)
		return nil, errors.New("更新用户失败")
	}

	return user, nil
}

func GetUser(id int64) (user *User, err error) {
	user = &User{Id: id}
	has, err := engine.Get(user)
	if has {
		return user, nil
	}
	return nil, errors.New("User not exists")
}

func GetUserByName(username string) (user *User, err error) {
	user = &User{Username: username}
	has, err := engine.Get(user)
	if has {
		return user, nil
	}
	return nil, errors.New("User not exists")
}
