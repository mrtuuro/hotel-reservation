package db

import (
	"context"
	"github.com/mrtuuro/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, bson.M, bson.M) error
	GetHotels(context.Context, bson.M) ([]*types.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("hotels"),
	}
}

func (m *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel, error) {
	cursor, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*types.Hotel
	if err = cursor.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (m *MongoHotelStore) Update(ctx context.Context, filter, update bson.M) error {
	_, err := m.coll.UpdateOne(ctx, filter, update)
	return err
}

func (m *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := m.coll.InsertOne(ctx, hotel, nil)
	if err != nil {
		return nil, err
	}

	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}
