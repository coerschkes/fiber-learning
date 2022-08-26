package userRoutes

import (
	userHandler "github.com/coerschkes/fiber-learning/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRouter struct {
	userHandler.UserHandler
}

func NewUserRouter(database *gorm.DB) *UserRouter {
	return &UserRouter{userHandler.NewUserHttpHandler(database)}
}

func (n UserRouter) SetupNoteRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Post("/", n.CreateUsers)
	user.Get("/", n.GetUsers)
	user.Get("/:noteId", n.GetUser)
	user.Put("/:noteId", n.UpdateUser)
	user.Delete("/:noteId", n.DeleteUser)
}
