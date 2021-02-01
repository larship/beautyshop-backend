package main

import (
	"github.com/larship/barbershop/config"
	"github.com/larship/barbershop/database"
	"github.com/larship/barbershop/server"
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
