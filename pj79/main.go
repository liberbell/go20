package main

import (
	"fmt"
	"pj79/apps/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test01"
	// u.Email = "test01@example.com"
	// u.Password = "testpassword"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test01rev2"
	// u.Email = "test01rev2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreateTodos("First Todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(2)
	// user.CreateTodos("Second Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	user, _ := models.GetUser(3)
	user.CreateTodos("Third Todo")

}
