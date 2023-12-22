package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
	
	func GetSourceCountry(c *fiber.Ctx) error {  
		// --------------------------------------------------------------   
			b 		:= new(models.SourceCountry)
		// --------------------------------------------------------------
			data, err := b.GetSourceCountry() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/sourcecountry/GetSourceCountry", fiber.Map{
				"Title": "SourceCountry",
				"Content": "SourceCountry",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 