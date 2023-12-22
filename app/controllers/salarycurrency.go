package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetSalaryCurrency(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.SalaryCurrency)
		// --------------------------------------------------------------
			data, err := b.GetSalaryCurrency() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/salarycurrency/GetSalaryCurrency", fiber.Map{
				"Title": "SalaryCurrency",
				"Content": "SalaryCurrency",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 