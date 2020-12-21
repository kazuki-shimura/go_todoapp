package main

import (
	"fmt"
	"log"
	"./config"
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
	u, _ := models.GetUser(2)
	fmt.Println(u)


	// Userの更新
	// u.Name = "Test2"
	// u.Email = "test2@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(2)
	// fmt.Println(u)


	// Userの削除
	// u.DeleteUser()
	// u, _ = models.GetUser(2)
	// fmt.Println(u)


	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")

	fmt.Println(models.Db)

	// People作成
	p := &models.People{}
	p.Firstname = "people2"
	p.Lastname = "people2"
	p.Email = "people2@gmail.com"
	p.Password = "people2"
	fmt.Println(p)
	p.CreatePeople()

}

