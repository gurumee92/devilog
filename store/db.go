package store

import (
	"fmt"
	"os"

	"github.com/gurumee92/devilog/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // init postgres
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // init sqlite
)

// GetDB is get db instance
func GetDB() *gorm.DB {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	name := os.Getenv("DATABASE_NAME")
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, name, pass)
	db, err := gorm.Open("postgres", url)

	if err != nil {
		fmt.Println("storage err: ", err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

// GetTestDB is get db instance for test
func GetTestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./test.db")

	if err != nil {
		fmt.Println("storage err: ", err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

// DropTestDB is drop test db
func DropTestDB() error {
	if err := os.Remove("./test.db"); err != nil {
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
