package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "0xroot/docs" // Swagger generated docs
	"0xroot/internal/config"
	"0xroot/internal/models"
	"0xroot/internal/services"
)

// @title My API
// @version 1.0
// @description This is my API
// @contact.name Amir
// @contact.email your-email@example.com
// @host localhost:3000
// @BasePath /api
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig("config.yaml")
		app := fiber.New()

		// Hello route
		// @Summary Hello endpoint
		// @Description Returns Hello message
		// @Tags example
		// @Success 200 {string} string "Hello world"
		// @Router / [get]
		app.Get("/", helloHandler)

		// Swagger
		app.Get("/swagger/*", fiberSwagger.WrapHandler)

		// API group
		api := app.Group("/api")
		auth := api.Group("/auth")

		// Register route
		auth.Post("/register", registerHandler)

		addr := fmt.Sprintf(":%d", cfg.Server.Port)
		app.Listen(addr)
	},
}

// helloHandler returns a simple hello message
func helloHandler(c *fiber.Ctx) error {
	return c.JSON("Hello world")
}

// registerHandler handles user registration
// @Summary Register a new user
// @Description Register a new user with email, password, and username
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User info"
// @Success 200 {object} map[string]string
// @Router /auth/register [post]
func registerHandler(c *fiber.Ctx) error {
	req := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}{}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	services.Register(user)

	return c.JSON(fiber.Map{
		"message": "Hello " + req.Email,
		"asdasd":  "asdasd",
	})
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
