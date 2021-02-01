package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/larship/barbershop/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	router *http.ServeMux
	config *config.Config
}

func Init(mux *http.ServeMux, conf *config.Config) *Server {
	return &Server{
		router: mux,
		config: conf,
	}
}

func (s *Server) Start() {
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	server := &http.Server{
		Addr:         s.config.WebServerAddress,
		Handler:      s.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	go gracefullyShutdown(server, quit, done)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Ошибка при создании веб-сервера: %s", err.Error())
	}

	<-done
	log.Printf("Веб-сервер остановлен")
}

func gracefullyShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	log.Println("Останавливаем веб-сервер...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Не смогли остановить веб-сервер: %v\n", err)
	}
	close(done)
}

func ResponseSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Ошибка преобразования объекта в JSON-строку: %v\n", err)
	}
}

func ResponseError(w http.ResponseWriter, statusCode int, errorText string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	data := map[string]string{
		"error": errorText,
	}
	dataStr, _ := json.Marshal(data)
	fmt.Fprintf(w, string(dataStr))
}
