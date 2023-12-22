package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetPostedDate(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.PostedDate)
		// --------------------------------------------------------------
			data, err := b.GetPostedDate() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/posteddate/GetPostedDate", fiber.Map{
				"Title": "PostedDate",
				"Content": "PostedDate",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 