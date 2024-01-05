package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func main() {
	// Load templates
	engine := html.New("./views", ".tmpl")

	// Creating a fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Configure app
	app.Static("/", "./public")

	// Add routes
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Name": "Mai",
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		var body struct {
			Message string
		}

		if err := c.BodyParser(&body); err != nil {
			return err
		}
		return c.Render("index", fiber.Map{
			"Name":    "Mai",
			"Message": body.Message,
		})
	})

	//Start app
	log.Fatal(app.Listen(":4000"))
}
