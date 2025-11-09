package main

import (
	"log"
	"os"

	"notes-backend/database"
	"notes-backend/middleware"
	"notes-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// load .env (optional)
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.Connect()

	app := fiber.New()

	// logger
	app.Use(middleware.Logger())
	app.Static("/uploads", "./uploads")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Authorization,Content-Type",
	}))
	// open routes: register/login
	routes.AuthRoutes(app)

	// protected group: require JWT
	// apply jwt middleware on /api routes
	app.Use("/api", middleware.RequireAuth())

	// notes endpoints
	routes.NoteRoutes(app)

	log.Printf("Listening on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
