package main

import (
	"app/db"
	"app/db/models"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&models.User{}, &models.LinkAbility{}, &models.MSCard{}, &models.PLCard{}, &models.TacticalCard{}, &models.GameDeck{})
}
