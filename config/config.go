package config

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	DatabaseDsn      string
	WebServerAddress string
	Smsc             smscConfig
}

type smscConfig struct {
	Login    string
	Password string
	Sender   string
}

func Init() *Config {
	configData := map[string]string{}
	for _, val := range os.Environ() {
		params := strings.Split(val, "=")
		configData[params[0]] = params[1]
	}

	conf := &Config{
		DatabaseDsn: "postgresql://" + configData["DATABASE_USER"] + ":" + configData["DATABASE_PASSWORD"] + "@" +
			configData["DATABASE_HOST"] + "/" + configData["DATABASE_NAME"],
		WebServerAddress: configData["WEB_SERVER_ADDRESS"],
		Smsc: smscConfig{
			Login:    configData["SMSC_LOGIN"],
			Password: configData["SMSC_PASSWORD"],
			Sender:   configData["SMSC_SENDER"],
		},
	}

	log.Printf("Конфиг: %+v", conf)

	return conf
}
