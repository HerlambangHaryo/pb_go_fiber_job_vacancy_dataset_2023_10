package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetJobFunctionValue(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.JobFunctionValue)
		// --------------------------------------------------------------
			data, err := b.GetJobFunctionValue() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/jobfunctionvalue/GetJobFunctionValue", fiber.Map{
				"Title": "JobFunctionValue",
				"Content": "JobFunctionValue",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 