package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtuuro/hotel-reservation/db"
	"github.com/mrtuuro/hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (u *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := u.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Jon",
		LastName:  "Waldo",
	}
	return c.JSON(user)
}
