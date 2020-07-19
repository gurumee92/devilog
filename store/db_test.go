package store

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/gurumee92/devilog/config"
	"github.com/gurumee92/devilog/model"
	"github.com/jinzhu/gorm"
)

var (
	c         *config.Config
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
	c = config.GetTestConfig()
	db = GetDB(c)
	AutoMigrate(db)
	postStore = NewPostStore(db)
	loadFixture()
}

func tearDown() {
	_ = db.Close()

	if err := DropTestDB(c); err != nil {
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
