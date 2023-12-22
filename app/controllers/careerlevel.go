package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetCareerLevel(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.CareerLevel)
		// --------------------------------------------------------------
			data, err := b.GetCareerLevel() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/careerlevel/GetCareerLevel", fiber.Map{
				"Title": "CareerLevel",
				"Content": "CareerLevel",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 