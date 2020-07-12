package store

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/gurumee92/devilog/model"
	"github.com/jinzhu/gorm"
)

var (
	db        *gorm.DB
	postStore *PostStore
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	db = GetTestDB()
	AutoMigrate(db)
	postStore = NewPostStore(db)

	loadFixture()
}

func tearDown() {
	_ = db.Close()

	if err := DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func loadFixture() {
	for i := 1; i <= 5; i++ {
		post := model.Post{
			Title:   "test" + strconv.Itoa(i),
			Content: "test" + strconv.Itoa(i),
			Author:  "test" + strconv.Itoa(i),
		}
		postStore.Save(&post)
	}
}
