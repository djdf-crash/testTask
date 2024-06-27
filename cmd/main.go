package main

import (
	"github.com/gofiber/fiber/v2/log"
	"testTask/cmd/internal"
	"testTask/cmd/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Panic("Error while read config: ", err)
		return
	}
	app, err := internal.NewApp(cfg)
	if err != nil {
		log.Panic("Error while create app: ", err)
		return
	}
	app.Run()
}
