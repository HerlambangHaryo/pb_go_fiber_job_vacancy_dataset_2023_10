package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models" 
		// "strconv" 
		"strings" 
	) 
	
	func GetJobtitle(c *fiber.Ctx) error { 
		// --------------------------------------------------------------   
			var persentase string
		// --------------------------------------------------------------   
			b 		:= new(models.Jobtitle)
		// --------------------------------------------------------------
			data, err := b.GetJobtitle() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			data2, err := b.GetJobtitlePercentage() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}  
			
			if len(data2) > 0 { 
				persentase 		= data2[0].Percentage 
			} else { 
				persentase 		= "" 
			}
		// --------------------------------------------------------------
			return c.Render("contents/jobtitle/GetJobtitle", fiber.Map{
				"Title": "Jobtitle",
				"Content": "Jobtitle",
				"Data"	: data,      
				"Data2"	: data2,  
				"persentase": persentase,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	
	func GetJobtitleWildcard(c *fiber.Ctx) error { 
		// --------------------------------------------------------------   
			name 			:= c.Params("val")
			replacedName 	:= strings.Replace(name, "_", " ", -1)  
		// --------------------------------------------------------------   
			b 		:= new(models.Jobtitle)
		// --------------------------------------------------------------
			data, err := b.GetJobtitleWildcard(replacedName) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/jobtitle/GetJobtitleWildcard", fiber.Map{
				"Title": "Jobtitle",
				"Content": "Jobtitle",
				"Data"	: data,       
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}