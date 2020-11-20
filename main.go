package main

import (
	"github.com/tomasarras/GoLang/controller/agencyController"
	"github.com/tomasarras/GoLang/database"
	"github.com/tomasarras/GoLang/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.StartConn()
	defer db.Close()

	router.Start()
	agencyController.Start(db)

}
