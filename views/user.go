package views

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"zy-o2o-service/models"
)

type UserView struct {
	Id       int64
	Username string
	Password string
	State    int8
	IsStaff  bool
	IsAdmin  bool
	LastTime time.Time
	Created  time.Time
	Updated  time.Time
}

var user = models.User{1,
	"admin",
	"1234qwer",
	1, true, true,
	time.Now(),
	time.Now(),
	time.Now(),
}

func GetUser(c *gin.Context) {
	user, err := models.GetUser(1)
	//user, err:=models.GetUserByName("admin")
	log.Println(user)
	log.Println(err)

}

func CreatedUser(c *gin.Context) {
	flag, err := models.AddUser(&user)
	log.Println(flag)
	log.Println(err)

}

func DeleteUser(c *gin.Context) {
	flag, err := models.DeleteUser(1)
	log.Println(flag)
	log.Println(err)

}
func UpdateUser(c *gin.Context) {
	user.Username = "2343"
	user, err := models.UpdateUser(&user)
	log.Println(user)
	log.Println(err)

}
