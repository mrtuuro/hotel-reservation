package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "Server's port")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiv1.Get("/user", handleUser)

	log.Fatal(app.Listen(*listenAddr))

}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working fine!"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "user here"})
}
