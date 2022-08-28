package main

import (
	"fmt"
	"log"
	"sample-server/cmd/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
) 



func main(){

	godotenv.Load()

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/health-check",healthCheck) 

	app.Post("/api/product",handlers.CreateProduct)

	fmt.Println("hello world")

	log.Fatal(app.Listen(":4500"))
}

func healthCheck(c*fiber.Ctx)error{
return c.SendString("Ok Boss")
}