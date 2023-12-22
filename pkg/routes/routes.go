package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.GetDashboard)  
   
	app.Get("/Jobtitle", controllers.GetJobtitle) 
	app.Get("/Jobtitle/Wildcard/:val", controllers.GetJobtitleWildcard) 

	app.Get("/DatasetJobtitle", controllers.GetDatasetJobtitle) 
	app.Get("/DatasetJobtitle/CreateId/:val", controllers.CreateDatasetJobtitleId) 
	app.Get("/DatasetJobtitle/CreateEn/:val", controllers.CreateDatasetJobtitleEn) 
	// app.Get("/DatasetJobtitle/CreateEn/:val", controllers.GetCreateDatasetJobtitleEn) 

	
	app.Get("/CareerLevel", controllers.GetCareerLevel) 
	app.Get("/DatasetCareerLevel", controllers.GetDatasetCareerLevel) 

	app.Get("/ClosingDate", controllers.GetClosingDate) 
	app.Get("/EmploymentType", controllers.GetEmploymentType) 
	app.Get("/FieldOfStudy", controllers.GetFieldOfStudy) 
	app.Get("/IndustryRequirement", controllers.GetIndustryRequirement) 
	app.Get("/IndustryValue", controllers.GetIndustryValue) 
	app.Get("/Location", controllers.GetLocation) 
	app.Get("/PostedDate", controllers.GetPostedDate) 
	app.Get("/Qualification", controllers.GetQualification) 
	app.Get("/SalaryCurrency", controllers.GetSalaryCurrency) 
	app.Get("/SalaryInvisible", controllers.GetSalaryInvisible) 
	app.Get("/SalaryMax", controllers.GetSalaryMax) 
	app.Get("/SalaryMin", controllers.GetSalaryMin) 
	app.Get("/SalaryType", controllers.GetSalaryType) 
	app.Get("/SourceCountry", controllers.GetSourceCountry) 
	app.Get("/YearsOfExperience", controllers.GetYearsOfExperience)  
	app.Get("/JobFunctionValue", controllers.GetJobFunctionValue)  

	
	app.Get("/Jobdescription/SetDone/:idx", controllers.GetJobdescriptionSetDone)  
}
