package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetClosingDate(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.ClosingDate)
		// --------------------------------------------------------------
			data, err := b.GetClosingDate() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/closingdate/GetClosingDate", fiber.Map{
				"Title": "ClosingDate",
				"Content": "ClosingDate",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 