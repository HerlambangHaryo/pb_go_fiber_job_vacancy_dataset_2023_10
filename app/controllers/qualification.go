package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetQualification(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.Qualification)
		// --------------------------------------------------------------
			data, err := b.GetQualification() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/qualification/GetQualification", fiber.Map{
				"Title": "Qualification",
				"Content": "Qualification",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 