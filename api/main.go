package main

import (
	"app/controller"
	"app/db"
	"app/db/models"
	libs "app/libs/firebase"
	repository "app/repositories"
	"app/router"
	usecase "app/usecase/commands"
	queryService "app/usecase/queries"
	"context"
	"fmt"
	"log"
)

func main() {
	dbConn := db.NewDB()
	ctx := context.Background()

	defer func() {
		db.CloseDB(dbConn)
		fmt.Println("Successfully Migrated and DB Connection Closed")
	}()

	// マイグレーションの実行
	err := dbConn.AutoMigrate(&models.User{}, &models.LinkAbility{}, &models.MSCard{}, &models.PLCard{}, &models.TacticalCard{}, &models.GameDeck{}, &models.SeriesTitle{}, &models.IncludeCode{})

	if err != nil {
		fmt.Println("マイグレーションに失敗しました:", err)
		return
	}

	firebaseService, err := libs.NewFirebaseService(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase service: %v", err)
	}

	userRepo := repository.NewUserRepository(dbConn)
	linkAbilityRepo := repository.NewLinkAbilityRepository(dbConn)
	plCardRepo := repository.NewPLCardRepository(dbConn)
	msCardRepo := repository.NewMSCardRepository(dbConn)
	tacRepo := repository.NewTacticalCardRepository(dbConn)
	includeCodeRepo := repository.NewIncludeCodeRepository(dbConn)
	seriesTitleRepo := repository.NewSeriesTitleRepository(dbConn)
	plCardQueryService := queryService.NewPlCardQueryService(dbConn)
	msCardQueryService := queryService.NewMsCardQueryService(dbConn)
	seriesTitleQueryService := queryService.NewSeriesTitleQueryService(dbConn)
	linkAbilityQueryService := queryService.NewLinkAbilityQueryService(dbConn)
	tacticalCardQueryService := queryService.NewTacticalCardQueryService(dbConn)
	includeCodeQueryService := queryService.NewIncludeCodeQueryService(dbConn)
	registerUsecase := usecase.NewRegisterUsecase(userRepo)
	importPlCardCsvUsecase := usecase.NewImportPlCardCsvUsecase(plCardRepo, linkAbilityRepo, includeCodeRepo)
	importMsCardCsvUsecase := usecase.NewImportMsCardCsvUsecase(msCardRepo, linkAbilityRepo, seriesTitleRepo, includeCodeRepo)
	importTacticalCardCsvUsecase := usecase.NewImportTacticalCardCsvUsecase(tacRepo)
	authController := controller.NewAuthController(registerUsecase, firebaseService)
	plCardController := controller.NewPlCardController(plCardQueryService, importPlCardCsvUsecase)
	msCardController := controller.NewMsCardController(importMsCardCsvUsecase, msCardQueryService)
	seriesTitleController := controller.NewSeriesTitleController(seriesTitleQueryService)
	linkAbilityController := controller.NewLinkAbilityController(linkAbilityQueryService)
	tacCardController := controller.NewTacticalCardController(importTacticalCardCsvUsecase, tacticalCardQueryService)
	includeCodeController := controller.NewIncludeCodeController(includeCodeQueryService)
	e := router.NewRouter(plCardController, msCardController, tacCardController, seriesTitleController, linkAbilityController, includeCodeController, authController)
	e.Logger.Fatal(e.Start(":8080"))
}
