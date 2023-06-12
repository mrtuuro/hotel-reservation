package db

import (
	"context"
	"github.com/mrtuuro/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		coll:       client.Database(DBNAME).Collection("rooms"),
		HotelStore: hotelStore,
	}
}

func (m *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := m.coll.InsertOne(ctx, room, nil)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room.ID}}
	if err = m.HotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}
	return room, nil
}

func (m *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	res, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*types.Room
	if err = res.All(ctx, &rooms); err != nil {
		return nil, err
	}
	
	return rooms, nil
}
