package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"test-api/club/config"
	"test-api/club/service/club"
)

func main() {
	//load Config
	c, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("[CONFIGURATION] Can not load configuration file.")
	}

	// Database init
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=%v",
		c.DbHost, c.DbUsername, c.DbPassword, c.DbName, c.DbPort, c.DbTz,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}, &gorm.Config{
		Logger: logger.Default.LogMode(config.GetDBLogLevel(c.DbLogLevel)),
	})
	if err != nil {
		log.Fatal("[DATABASE] Database connection failed.")
	}
	log.Println("[DATABASE] Database connection success.")

	//Auto Migrate Database
	log.Println("Registering table..")
	err = config.Club(db)
	if err != nil {
		log.Fatal("[DATABASE] Database staff cant migrate")
	}
	log.Println("[DATABASE] Database migrate success.")

	router := config.SetupRouter(c)
	clubRouter := router.Group("/club")

	club.RouteClub(db, clubRouter)

	router.Run()
}
