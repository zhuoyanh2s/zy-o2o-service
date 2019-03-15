package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"zy-o2o-service/models"
	"zy-o2o-service/routers"
)

func main() {
	viper.SetConfigName("dev")      // name of config file (without extension)
	viper.AddConfigPath("./config") // path to look for the config file in
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	models.Setup()

	routers.Setup()
	gin.SetMode(viper.GetString("server.RunMode"))
	routers.Engine.Run()
}
