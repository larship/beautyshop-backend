package main

import (
	"github.com/larship/beautyshop/config"
	"github.com/larship/beautyshop/database"
	"github.com/larship/beautyshop/server"
	"net/http"
)

func main() {
	conf := config.Init()
	db := database.Init(conf)

	mux := http.NewServeMux()
	s := server.Init(mux, conf)
	s.MakeRoutes()
	s.Start()

	db.CloseConnection()
}
