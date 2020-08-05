package main

import (
	"github.com/gurumee92/devilog/config"
	"github.com/gurumee92/devilog/handler"
	"github.com/gurumee92/devilog/router"
	"github.com/gurumee92/devilog/store"
)

func main() {
	c := config.GetConfig()
	// database
	db := store.GetDB(c)
	defer db.Close()
	store.AutoMigrate(db)
	postStore := store.NewPostStore(db)
	accountStore := store.NewAccountStore(db)

	// echo
	e := router.NewRouter(c)
	h := handler.NewHandler(c, postStore, accountStore)
	h.Register(e)
	e.Logger.Fatal(e.Start(":8080"))
}
