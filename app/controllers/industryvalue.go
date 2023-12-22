package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetIndustryValue(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.IndustryValue)
		// --------------------------------------------------------------
			data, err := b.GetIndustryValue() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/industryvalue/GetIndustryValue", fiber.Map{
				"Title": "IndustryValue",
				"Content": "IndustryValue",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 