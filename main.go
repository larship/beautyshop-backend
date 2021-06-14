package main

import (
	"github.com/larship/beautyshop/config"
	"github.com/larship/beautyshop/database"
	"github.com/larship/beautyshop/server"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	log.Print("Запускаем приложение")

	rand.Seed(time.Now().UnixNano())

	conf := config.Init()
	db := database.Init(conf)

	mux := http.NewServeMux()
	s := server.Init(mux, conf)
	s.MakeRoutes()
	s.Start()

	db.CloseConnection()
}
