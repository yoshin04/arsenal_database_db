package main

import (
	"app/controller"
	"app/db"
	"app/db/models"
	repository "app/repositories"
	"app/router"
	usecase "app/usecase/commands"
	queryService "app/usecase/queries"
	"fmt"
)

func main() {
	dbConn := db.NewDB()

	defer func() {
		db.CloseDB(dbConn)
		fmt.Println("Successfully Migrated and DB Connection Closed")
	}()

	// マイグレーションの実行
	err := dbConn.AutoMigrate(&models.User{}, &models.LinkAbility{}, &models.MSCard{}, &models.PLCard{}, &models.TacticalCard{}, &models.GameDeck{})
	if err != nil {
		fmt.Println("マイグレーションに失敗しました:", err)
		return
	}

	linkAbilityRepo := repository.NewLinkAbilityRepository(dbConn)
	plCardRepo := repository.NewPLCardRepository(dbConn)
	plCardQueryService := queryService.NewPlCardQueryService()
	importPlCardCsvUsecase := usecase.NewImportPlCardCsvUsecase(plCardRepo, linkAbilityRepo)
	plCardController := controller.NewPlCardController(plCardQueryService, importPlCardCsvUsecase)
	e := router.NewRouter(plCardController)
	e.Logger.Fatal(e.Start(":8080"))
}
