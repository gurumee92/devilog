package main

import (
	"github.com/gurumee92/devilog/handler"
	"github.com/gurumee92/devilog/router"
	"github.com/gurumee92/devilog/store"
)

func main() {
	// database
	db := store.GetDB()
	defer db.Close()
	store.AutoMigrate(db)
	postStore := store.NewPostStore(db)

	// echo
	e := router.NewRouter()
	h := handler.NewHandler(postStore)
	h.Register(e)
	e.Logger.Fatal(e.Start(":1323"))
}
