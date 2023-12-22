package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetSalaryType(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.SalaryType)
		// --------------------------------------------------------------
			data, err := b.GetSalaryType() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/salarytype/GetSalaryType", fiber.Map{
				"Title": "SalaryType",
				"Content": "SalaryType",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 