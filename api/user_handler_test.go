package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtuuro/hotel-reservation/db"
	"github.com/mrtuuro/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http/httptest"
	"testing"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) Teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.Background()); err != nil {
		t.Fatal()
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		db.NewMongoUserStore(client, db.TestDBNAME),
	}
}

func TestPostUser(t *testing.T) {
	testDb := setup(t)
	defer testDb.Teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testDb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		FirstName: "tunc",
		LastName:  "ozay",
		Email:     "tunc@tunc.com",
		Password:  "1234567",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	response, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var user types.User
	json.NewDecoder(response.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("expecting a user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expecting the EncryptedPassword not to be included in the json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected last name %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}
}
