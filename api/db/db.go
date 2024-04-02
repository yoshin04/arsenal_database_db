package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	var db *gorm.DB
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	// リトライロジック
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("データベースに接続しました")
			return db
		}
		fmt.Println("データベースへの接続に失敗しました。リトライします...")
		time.Sleep(5 * time.Second) // 5秒待機
	}

	log.Panic("データベースへの接続に最終的に失敗しました: ", err, "dsn: ", dsn)
	return nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("データベースの切断に失敗しました")
	}

	sqlDB.Close()
	fmt.Println("データベースを切断しました")
}
