package database

import (
	"dev-oleksandrv/taskera-app/internal/config"
	"dev-oleksandrv/taskera-app/internal/model/domain"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	dbCfg := config.GetConfig().DBConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Database,
		dbCfg.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}

	DB = db
	log.Println("Successfully connected to database")
}
