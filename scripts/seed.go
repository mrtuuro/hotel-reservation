package main

import (
	"context"
	"github.com/mrtuuro/hotel-reservation/db"
	"github.com/mrtuuro/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func seedHotel(name, location string, rating int) {
	hotel := &types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []types.Room{
		{
			Type:      types.SeasideRoomType,
			BasePrice: 56.4,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 123.54,
		},
		{
			Type:      types.DeluxeRoomType,
			BasePrice: 76.43,
		},
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		_, err = roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	seedHotel("Continental", "NYC", 6)
	seedHotel("Bellucia", "France", 7)
	seedHotel("Hades", "London", 5)

}

func init() {
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
}
