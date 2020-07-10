package main

import (
	"github.com/gurumee92/devilog/handler"
	"github.com/gurumee92/devilog/router"
)

func main() {
	e := router.NewRouter()
	h := handler.NewHandler()
	h.Register(e)
	e.Logger.Fatal(e.Start(":1323"))
}
