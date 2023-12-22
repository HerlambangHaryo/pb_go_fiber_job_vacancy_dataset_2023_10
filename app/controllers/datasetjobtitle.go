package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models" 
		// "strconv" 
		"strings"  
	) 
  
	func GetDatasetJobtitle(c *fiber.Ctx) error {
		// --------------------------------------------------------------   
			b 		:= new(models.DatasetJobtitle)
		// --------------------------------------------------------------
			data, err := b.GetDatasetJobtitle() 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			return c.Render("contents/datasetjobtitle/GetDatasetJobtitle", fiber.Map{
				"Title": "Dataset Jobtitle",
				"Content": "DatasetJobtitle",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func CreateDatasetJobtitleId(c *fiber.Ctx) error {
		// --------------------------------------------------------------   
			var idx uint
		// --------------------------------------------------------------     
			name 			:= c.Params("val")
			replacedName 	:= strings.Replace(name, "_", " ", -1) 
		// --------------------------------------------------------------    
			djt := &models.DatasetJobtitle{}
			jt 	:= &models.Jobtitle{}
		// --------------------------------------------------------------   
			err := djt.InsertNama(replacedName)
			
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}   
		// --------------------------------------------------------------
			data2, err := djt.GetDatasetJobtitleByNameId(replacedName) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}  
			
			if len(data2) > 0 { 
				idx 		= data2[0].ID 
			} else { 
				idx 		= 0 
			}
		// --------------------------------------------------------------   
			// // UPDATE `jobtitle` SET  `jobtitle_id`='?  WHERE name = ?  
			// data3, err := jt.UpdateJobtitle(idx int, replacedName) 
			data3 := jt.UpdateJobtitle(replacedName, "jobtitle_id", idx)

			if data3 != nil {
				return c.Status(500).SendString(data3.Error())
			} 
		// --------------------------------------------------------------   
			// // UPDATE `dataset_jobtitle` SET  `equal`= ?  WHERE id = ?
			data4 := djt.UpdateDatasetJobtitle(idx, "sama_dengan", 1)
			// err := djt.UpdateDatasetJobtitle(jobtitle_id int, replacedName string) 

			if data4 != nil {
				return c.Status(500).SendString(data4.Error())
			}
		// --------------------------------------------------------------  
			previousPage := c.Get("Referer")
			if previousPage == "" {
				previousPage = "/default/page"
			}

			return c.Redirect(previousPage)
		// --------------------------------------------------------------
	}

	func CreateDatasetJobtitleEn(c *fiber.Ctx) error {
		// --------------------------------------------------------------   
			var idx uint
		// --------------------------------------------------------------    
			name 			:= c.Params("val")
			replacedName 	:= strings.Replace(name, "_", " ", -1) 
		// --------------------------------------------------------------    
			djt := &models.DatasetJobtitle{}
			jt 	:= &models.Jobtitle{}
		// --------------------------------------------------------------   
			err := djt.InsertName(replacedName)
			
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}   
		// --------------------------------------------------------------
			data2, err := djt.GetDatasetJobtitleByNameEn(replacedName) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}  
			
			if len(data2) > 0 { 
				idx 		= data2[0].ID 
			} else { 
				idx 		= 0 
			}
		// --------------------------------------------------------------   
			// // UPDATE `jobtitle` SET  `jobtitle_id`='?  WHERE name = ?  
			// data3, err := jt.UpdateJobtitle(idx int, replacedName) 
			data3 := jt.UpdateJobtitle(replacedName, "jobtitle_id", idx)

			if data3 != nil {
				return c.Status(500).SendString(data3.Error())
			} 
		// --------------------------------------------------------------   
			// // UPDATE `dataset_jobtitle` SET  `equal`= ?  WHERE id = ?
			data4 := djt.UpdateDatasetJobtitle(idx, "equal", 1)
			// err := djt.UpdateDatasetJobtitle(jobtitle_id int, replacedName string) 

			if data4 != nil {
				return c.Status(500).SendString(data4.Error())
			}
		// --------------------------------------------------------------  
			previousPage := c.Get("Referer")
			if previousPage == "" {
				previousPage = "/default/page"
			} 

			return c.Redirect(previousPage)
		// --------------------------------------------------------------
	}
 
 
 
 