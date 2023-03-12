package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"object-storage.app/buckets"
	"object-storage.app/handlers"
	"object-storage.app/object_creator"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {

	// database.StartMySql()
	object_creator.WriteCheeseMapToJSONFile()

	buckets.CreateBucket("object-storage-railway", "food-bucket-dev")
	buckets.WriteToBucket("food-bucket-dev", "cheese.json")
	buckets.ReadFromBucket("food-bucket-dev", "2023_03_12_cheese.json")
	object_creator.DeleteJsonFile()

	app := fiber.New()

	app.Get("/api/v1/health", handlers.HandlerHealthCheck)

	app.Listen(getPort())

}
