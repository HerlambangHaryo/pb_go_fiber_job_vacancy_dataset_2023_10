package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 
  
	func GetDatasetCareerLevel(c *fiber.Ctx) error {
		// --------------------------------------------------------------   
			b 		:= new(models.DatasetCareerLevel)
		// --------------------------------------------------------------
			data, err := b.GetDatasetCareerLevel() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			return c.Render("contents/datasetcareerlevel/GetDatasetCareerLevel", fiber.Map{
				"Title": "Dataset CareerLevel",
				"Content": "DatasetCareerLevel",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}