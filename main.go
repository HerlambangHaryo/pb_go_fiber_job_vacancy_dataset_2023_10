// main.go
package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database"
	// "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models" 
	"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/pkg/routes"
	"github.com/gofiber/template/html/v2"

)

func main() {
	connectToDatabase()

	engine := html.New("./resources/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

    app.Static("/", "./public")
   
	app.Use(logger.New())
	registerRoutes(app)
	log.Fatal(app.Listen(":8001"))
}

func connectToDatabase() {
	database.Connect() 
}

func registerRoutes(app *fiber.App) {
	routes.SetupRoutes(app)
}
