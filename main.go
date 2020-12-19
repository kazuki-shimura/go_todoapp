package main

import (
	"fmt"
	// "log"
	// "./config"
	"./models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// User作成
	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@gmail.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()


	// Userの取得
	u, _ := models.GetUser(1)
	fmt.Println(u)

	u.Name = "Test2"
	u.Email = "test2@gmail.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)
}

