package main

import (
	"fmt"
	// "log"
	// "./config"
	"./models"
	// "./controllers"
)

// func main() {
// 	fmt.Println(models.Db)
// 	controllers.StartMainServer()

// }

// func main() {
// 	// fmt.Println(models.Db)
// 	// fmt.Println(config.Config.Port)
// 	// fmt.Println(config.Config.SQLDriver)
// 	// fmt.Println(config.Config.DbName)
// 	// fmt.Println(config.Config.LogFile)

// 	// log.Println("test")

// 	// fmt.Println(models.Db)

// 	// User作成
// 	// u := &models.User{}
// 	// u.Name = "test"
// 	// u.Email = "test@gmail.com"
// 	// u.PassWord = "testtest"
// 	// fmt.Println(u)
// 	// u.CreateUser()


// 	// Userの取得
// 	// u, _ := models.GetUser(2)
// 	// fmt.Println(u)


// 	// Userの更新
// 	// u.Name = "Test2"
// 	// u.Email = "test2@gmail.com"
// 	// u.UpdateUser()
// 	// u, _ = models.GetUser(2)
// 	// fmt.Println(u)


// 	// Userの削除
// 	// u.DeleteUser()
// 	// u, _ = models.GetUser(2)
// 	// fmt.Println(u)



// 	// -------------------------------------------
// 	// fmt.Println(config.Config.Port)
// 	// fmt.Println(config.Config.SQLDriver)
// 	// fmt.Println(config.Config.DbName)
// 	// fmt.Println(config.Config.LogFile)

// 	// log.Println("test")

// 	// fmt.Println(models.Db)


// 	// People作成
// 	// p := &models.People{}
// 	// p.Firstname = "people2"
// 	// p.Lastname = "people2"
// 	// p.Email = "people2@gmail.com"
// 	// p.Password = "people2"
// 	// fmt.Println(p)
// 	// p.CreatePeople()

// 	// p, _ := models.GetPeople(1)
// 	// fmt.Println(p)

// peoples, _ := models.GetAllPeople()
// 	for _, v := range peoples {
// 		fmt.Println(v)
// 	}
// 	fmt.Println("-------------------------")

// 	// Peopleの更新
// 	// p.Firstname = "Test2People"
// 	// p.Lastname = "22222"
// 	// p.Email = "test2People@gmail.com"
// 	// p.UpdatePeople()
// 	// p, _ = models.GetPeople(1)
// 	// fmt.Println(p)

// 	// Peopleの削除
// 	// p.DeletePeople()
// 	// p, _ = models.GetPeople(1)
// 	// fmt.Println(p)

// 	// user, _ := models.GetUser(3)
// 	// user.CreateTodo("5 tfdafdafafodo")

// 	// t, _ := models.GetTodo(1)
// 	// fmt.Println(t)

// 	todos, _ := models.GetTodos()
// 	for _, v := range todos {
// 		fmt.Println(v)
// 	}
// 	fmt.Println("-------------------------")

// 	user3, _ := models.GetUser(3)
// 	todos3, _ := user3.GetTodosByUser()
// 	for _, v := range todos3 {
// 		fmt.Println(v)
// 	}

// 	// t, _ := models.GetTodo(1)
// 	// t.Content = "Update Todo"
// 	// t.UpdateTodo()

// 	// t, _ := models.GetTodo(9)
// 	// t.DeleteTodo()

// }

func main() {
	// fmt.Println(models.Db)
	// AAAの作成
	// a := &models.AAA{}
	// a.First = "firstAAA"
	// a.Second = "secondAAA"
	// a.Third = "thirdAAA"
	// a.Email = "aaa@gmail.com"
	// a.Password = "aaaaaaaa"
	// fmt.Println(a)
	// a.CreateAAA()

	a, _ := models.GetAAA(1)
	fmt.Println(a)
	fmt.Println("-------------------------")
	aaas, _ := models.GetAAAs()
	for _, v := range aaas {
		fmt.Println(v)
	}
}

