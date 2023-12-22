package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models"  
	) 

	// func GetJobdescriptionSetDone(c *fiber.Ctx) error { 
	// 	// --------------------------------------------------------------     
	// 		idx 			:= c.Params("idx") 
	// 	// --------------------------------------------------------------     
	// 		jt 	:= &models.Jobdescription{}
	// 	// --------------------------------------------------------------     
	// 		data4 := jt.GetJobdescriptionSetDone(idx, "done", 1) 

	// 		if data4 != nil {
	// 			return c.Status(500).SendString(data4.Error())
	// 		}
	// 	// --------------------------------------------------------------   
	// 		// Define the URL to redirect to
	// 		page := "/Dashboard"
        
	// 		// Redirect to the specified page
	// 		return c.Redirect(page, fiber.StatusFound) 
	// 	// --------------------------------------------------------------
	// }
	 

	
	func GetJobdescriptionSetDone(c *fiber.Ctx) error {
		idx := c.Params("idx")

		jt := &models.Jobdescription{}

		err := jt.GetJobdescriptionSetDone(idx, "done", 1)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		// Define the URL to redirect to
		page := "/Dashboard/" +  idx

		// Redirect to the specified page
		return c.Redirect(page, fiber.StatusFound)
	}