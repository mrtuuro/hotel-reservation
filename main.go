package main

import (
	"context"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtuuro/hotel-reservation/api"
	"github.com/mrtuuro/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const dburi = "mongodb://localhost:27017"

func main() {
	var config = fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.JSON(map[string]string{"error": err.Error()})
		},
	}

	listenAddr := flag.String("listenAddr", ":3000", "Server's port")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	// handler initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	v1 := app.Group("/api/v1")

	v1.Post("/user", userHandler.HandlePostUser)
	v1.Get("/user", userHandler.HandleGetUsers)
	v1.Get("/user/:id", userHandler.HandleGetUser)

	log.Fatal(app.Listen(*listenAddr))

}
