package main

import (
	"app/controller"
	"app/db"
	"app/router"
	queryService "app/usecase/queries"
)

func main() {
	db.NewDB()
	plCardQueryService := queryService.NewPlCardQueryService()
	plCardController := controller.NewPlCardController(plCardQueryService)
	e := router.NewRouter(plCardController)
	e.Logger.Fatal(e.Start(":8080"))
}
