package store

import (
	"log"
	"os"

	"github.com/gurumee92/devilog/config"
	"github.com/gurumee92/devilog/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // init postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // init sqlite
)

// GetDB is get db instance
func GetDB(c *config.Config) *gorm.DB {
	db, err := gorm.Open(c.DatabaseDialect, c.DatabaseURL)

	if err != nil {
		log.Fatalln("storage err: ", c.DatabaseURL)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(c.IsProduct)
	// db.LogMode(!c.IsProduct)
	return db
}

// DropTestDB is drop test db
func DropTestDB(c *config.Config) error {
	if err := os.Remove(c.DatabaseURL); err != nil {
		return err
	}
	return nil
}

// AutoMigrate is migrate code and database
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Post{},
	)
}
