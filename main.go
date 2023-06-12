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

func main() {
	var config = fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.JSON(map[string]string{"error": err.Error()})
		},
	}

	listenAddr := flag.String("listenAddr", ":3000", "Server's port")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// handler initialization
	var (
		hotelStore = db.NewMongoHotelStore(client)
		roomStore  = db.NewMongoRoomStore(client, hotelStore)
		userStore  = db.NewMongoUserStore(client)
		store      = &db.Store{
			User:  userStore,
			Hotel: hotelStore,
			Room:  roomStore,
		}
		userHandler  = api.NewUserHandler(userStore)
		hotelHandler = api.NewHotelHandler(store)
		app          = fiber.New(config)
		v1           = app.Group("/api/v1")
	)
	// User handlers
	v1.Post("/user", userHandler.HandlePostUser)
	v1.Get("/user", userHandler.HandleGetUsers)
	v1.Get("/user/:id", userHandler.HandleGetUser)
	v1.Put("/user/:id", userHandler.HandlePutUser)
	v1.Delete("/user/:id", userHandler.HandleDeleteUser)

	// Hotel handlers
	v1.Get("/hotel", hotelHandler.HandleGetHotels)
	v1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	v1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)
	log.Fatal(app.Listen(*listenAddr))

}
