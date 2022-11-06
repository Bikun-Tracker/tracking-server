package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"tracking-server/di"
	"tracking-server/shared/config"
)

func main() {
	container := di.Container

	err := container.Invoke(func(http *fiber.App, env *config.EnvConfig) error {
		log.Println(env)
		err := http.Listen(":" + env.PORT)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("error when starting http server: %s", err.Error())
	}
}
