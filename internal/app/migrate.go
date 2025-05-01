package app

import (
	"fmt"
	"log"

	"github.com/vnFuhung2903/rubyams-backend/config"
	"github.com/vnFuhung2903/rubyams-backend/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateDatabase(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	err = db.AutoMigrate(&entity.Student{}, &entity.Bid{}, &entity.Vote{})
	if err != nil {
		log.Fatalf("Migrate error: %v", err)
	}
	return db
}
