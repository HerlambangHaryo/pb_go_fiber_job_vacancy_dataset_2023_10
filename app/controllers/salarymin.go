package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetSalaryMin(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.SalaryMin)
		// --------------------------------------------------------------
			data, err := b.GetSalaryMin() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/salarymin/GetSalaryMin", fiber.Map{
				"Title": "SalaryMin",
				"Content": "SalaryMin",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 