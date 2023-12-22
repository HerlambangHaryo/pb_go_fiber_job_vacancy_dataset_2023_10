package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetApplyUrl(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.ApplyUrl)
		// --------------------------------------------------------------
			data, err := b.GetApplyUrl() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/applyurl/GetApplyUrl", fiber.Map{
				"Title": "ApplyUrl",
				"Content": "ApplyUrl",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 