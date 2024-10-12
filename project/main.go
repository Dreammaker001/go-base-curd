package main

import (
	"time"

	v1 "github.com/Dreammaker001/go-base-curd/route/Api/V1"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true, // true, run several processes so that all CPUs are used optimally
	})

	route := v1.Route(app)

	err := route.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
