package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtuuro/hotel-reservation/api"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "Server's port")
	flag.Parse()
	app := fiber.New()
	v1 := app.Group("/api/v1")

	v1.Get("/user", api.HandleGetUsers)
	v1.Get("/user/:id", api.HandleGetUserById)

	log.Fatal(app.Listen(*listenAddr))

}
