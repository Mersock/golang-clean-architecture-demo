package main

import (
	"fmt"

	"github.com/Mersock/golang-clean-architecture-demo/databases"
	"github.com/Mersock/golang-clean-architecture-demo/deliveries/routes"
	"github.com/Mersock/golang-clean-architecture-demo/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	databases.DB, err = gorm.Open("mysql", databases.DbURL(databases.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer databases.DB.Close()
	// run the migrations: todo struct
	databases.DB.AutoMigrate(&models.Todo{})
	//setup routes
	r := routes.SetupRouter()
	// running
	r.Run(":8080")
}
