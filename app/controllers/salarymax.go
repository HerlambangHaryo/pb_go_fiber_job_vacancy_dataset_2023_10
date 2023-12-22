package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetSalaryMax(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.SalaryMax)
		// --------------------------------------------------------------
			data, err := b.GetSalaryMax() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/salarymax/GetSalaryMax", fiber.Map{
				"Title": "SalaryMax",
				"Content": "SalaryMax",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 