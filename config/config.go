package config

import (
	"flag"
)

type Config struct {
	DatabaseDsn      string
	WebServerAddress string
}

func Init() *Config {
	// @todo Заполнять значения по-умолчанию на основе переменных окружения (environment variables)
	// Либо передавать во флаги переменные окружения, например, go run . -web-server-address="$BARBERSHOP_WEB_SERVER_ADDRESS"
	databaseDsn := flag.String("database-dsn", "postgresql://barbershop:barbershop456498@localhost:5432/barbershop", "DSN для подключения к БД")
	webServerAddress := flag.String("web-server-address", ":8080", "Адрес, который будет слушать веб-сервер")
	flag.Parse()

	conf := &Config{
		DatabaseDsn:      *databaseDsn,
		WebServerAddress: *webServerAddress,
	}

	return conf
}
