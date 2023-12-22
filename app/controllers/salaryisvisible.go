package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetSalaryInvisible(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.SalaryInvisible)
		// --------------------------------------------------------------
			data, err := b.GetSalaryInvisible() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/salaryinvisible/GetSalaryInvisible", fiber.Map{
				"Title": "SalaryInvisible",
				"Content": "SalaryInvisible",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 