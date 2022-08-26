package userHandler

import (
	"github.com/coerschkes/fiber-learning/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler interface {
	GetUsers(c *fiber.Ctx) error
	CreateUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
}

type UserHttpHandler struct {
	database *gorm.DB
}

func NewUserHttpHandler(database *gorm.DB) *UserHttpHandler {
	return &UserHttpHandler{database}
}

func (h UserHttpHandler) GetUsers(c *fiber.Ctx) error {
	var users []model.User

	h.database.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No users present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": users})
}

func (h UserHttpHandler) CreateUsers(c *fiber.Ctx) error {
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	user.ID = uuid.New()
	err = h.database.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func (h UserHttpHandler) GetUser(c *fiber.Ctx) error {
	var user model.User

	id := c.Params("userId")

	h.database.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Users Found", "data": user})
}

func (h UserHttpHandler) UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Name     string `json:"name"`
		Forename string `json:"forename"`
	}
	var user model.User

	id := c.Params("userId")

	h.database.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user.Name = updateUserData.Name
	user.Forename = updateUserData.Forename

	h.database.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": user})
}

func (h UserHttpHandler) DeleteUser(c *fiber.Ctx) error {
	var user model.User

	id := c.Params("userId")

	h.database.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user present", "data": nil})
	}

	err := h.database.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Deleted User"})
}
