package main

import (
	"app/controller"
	"app/router"
	queryService "app/usecase/queries"
)

func main() {
	plCardQueryService := queryService.NewPlCardQueryService()
	plCardController := controller.NewPlCardController(plCardQueryService)
	e := router.NewRouter(plCardController)
	e.Logger.Fatal(e.Start(":8080"))
}
