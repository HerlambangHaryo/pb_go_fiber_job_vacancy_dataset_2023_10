package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.GetDashboard)  
   
	app.Get("/Jobtitle", controllers.GetJobtitle) 
	app.Get("/Jobtitle/Wildcard/:val", controllers.GetJobtitleWildcard) 

	app.Get("/DatasetJobtitle", controllers.GetDatasetJobtitle) 
	app.Get("/DatasetJobtitle/CreateId/:val", controllers.CreateDatasetJobtitleId) 
	app.Get("/DatasetJobtitle/CreateEn/:val", controllers.CreateDatasetJobtitleEn) 
	// app.Get("/DatasetJobtitle/CreateEn/:val", controllers.GetCreateDatasetJobtitleEn) 
}
