package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetEmploymentType(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.EmploymentType)
		// --------------------------------------------------------------
			data, err := b.GetEmploymentType() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/employmenttype/GetEmploymentType", fiber.Map{
				"Title": "EmploymentType",
				"Content": "EmploymentType",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 