package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetFieldOfStudy(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.FieldOfStudy)
		// --------------------------------------------------------------
			data, err := b.GetFieldOfStudy() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/fieldofstudy/GetFieldOfStudy", fiber.Map{
				"Title": "FieldOfStudy",
				"Content": "FieldOfStudy",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 