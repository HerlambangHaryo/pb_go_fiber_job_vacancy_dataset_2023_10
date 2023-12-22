package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetYearsOfExperience(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.YearsOfExperience)
		// --------------------------------------------------------------
			data, err := b.GetYearsOfExperience() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/yearsofexperience/GetYearsOfExperience", fiber.Map{
				"Title": "YearsOfExperience",
				"Content": "YearsOfExperience",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 