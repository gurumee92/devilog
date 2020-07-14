package handler

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/gurumee92/devilog/config"
	"github.com/gurumee92/devilog/model"
	"github.com/gurumee92/devilog/router"
	"github.com/gurumee92/devilog/store"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

var (
	c         *config.Config
	h         *Handler
	e         *echo.Echo
	db        *gorm.DB
	postStore *store.PostStore
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	c = config.GetTestConfig()
	db = store.GetDB(c)
	store.AutoMigrate(db)
	postStore = store.NewPostStore(db)
	loadFixture()

	e = router.NewRouter()
	h = NewHandler(postStore)
	h.Register(e)
}

func tearDown() {
	_ = db.Close()

	if err := store.DropTestDB(c); err != nil {
		log.Fatal(err)
	}

	h = nil
	e = nil
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
