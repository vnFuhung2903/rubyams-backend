package config

import (
	"fmt"
	"log"

	"github.com/vnFuhung2903/rubyams-backend/env"
	"github.com/vnFuhung2903/rubyams-backend/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresDb(env *env.Env) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", env.DBHost, env.DBUser, env.DBPass, env.DBName, env.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}

	err = db.AutoMigrate(&model.Student{}, &model.Bid{}, &model.Vote{})
	if err != nil {
		log.Fatalf("Migrate error: %v", err)
	}
	return db
}
