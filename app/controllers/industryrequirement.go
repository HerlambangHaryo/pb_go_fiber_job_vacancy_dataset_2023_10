package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetIndustryRequirement(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.IndustryRequirement)
		// --------------------------------------------------------------
			data, err := b.GetIndustryRequirement() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/industryrequirement/GetIndustryRequirement", fiber.Map{
				"Title": "IndustryRequirement",
				"Content": "IndustryRequirement",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 