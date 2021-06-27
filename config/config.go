package config

import (
	"bufio"
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
	file, err := os.Open(".env")

	if err != nil {
		log.Printf("Ошибка при создании веб-сервера: %s", err.Error())
		return nil
	}

	defer file.Close()

	configData := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		params := strings.Split(scanner.Text(), "=")
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
