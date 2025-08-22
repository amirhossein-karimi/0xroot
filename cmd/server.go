package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/spf13/cobra"

	"0xroot/internal/config"
	"0xroot/internal/models"
	"0xroot/internal/services"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadConfig("config.yaml")
		app := fiber.New()

		app.Get("/", func(c fiber.Ctx) error {
			return c.JSON("Hello wolrd")
		})

		api := app.Group("/api")

		auth := api.Group("/auth")

		auth.Post("/register", func(c fiber.Ctx) error {

			req := struct {
				Email    string `json:"email"`
				Password string `json:"password"`
				Username string `json:"username"`
			}{}

			if err := c.Bind().JSON(&req); err != nil {
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
			})
		})

		addr := fmt.Sprintf(":%d", cfg.Server.Port)
		app.Listen(addr)
		app.Listen(":3000")
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
