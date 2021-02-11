package main

import (
	"github.com/larship/beautyshop/config"
	"github.com/larship/beautyshop/database"
	"github.com/larship/beautyshop/server"
	"log"
	"net/http"
)

func main() {
	log.Print("Запускаем приложение")

	conf := config.Init()
	db := database.Init(conf)

	mux := http.NewServeMux()
	s := server.Init(mux, conf)
	s.MakeRoutes()
	s.Start()

	db.CloseConnection()
}
