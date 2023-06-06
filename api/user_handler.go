package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrtuuro/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Jon",
		LastName:  "Waldo",
	}
	return c.JSON(user)
}

func HandleGetUserById(c *fiber.Ctx) error {
	return c.JSON("one user returned")
}
