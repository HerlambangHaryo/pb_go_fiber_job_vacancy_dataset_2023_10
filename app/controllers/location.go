package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetLocation(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.Location)
		// --------------------------------------------------------------
			data, err := b.GetLocation() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/location/GetLocation", fiber.Map{
				"Title": "Location",
				"Content": "Location",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 