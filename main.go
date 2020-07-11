package main

import (
	"github.com/gurumee92/devilog/handler"
	"github.com/gurumee92/devilog/router"
	"github.com/gurumee92/devilog/store"
)

func main() {
	e := router.NewRouter()
	db := store.GetDB()
	defer db.Close()
	store.AutoMigrate(db)
	// postStore := store.NewPostStore(db)

	h := handler.NewHandler()
	h.Register(e)
	e.Logger.Fatal(e.Start(":1323"))
}
