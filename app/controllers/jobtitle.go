package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/app/models" 
		"strconv" 
	) 
	
	func GetFixtureInformation(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			leagueapi_id := c.Params("leagueapi_id")
			season := c.Params("season")
			fixtureapi_id := c.Params("fixtureapi_id")
	
			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season)
			fixtureapi_idInt, err 	:= strconv.Atoi(fixtureapi_id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			b 		:= new(models.FootballFixture)
		// --------------------------------------------------------------
			data, err := b.GetFixtureGeneral(uint(leagueapi_idInt), 
						uint(seasonInt), 
						uint(fixtureapi_idInt)) 

			if err != nil {
			return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			return c.Render("contents/fixtures/GetFixtureInformation", fiber.Map{
				"Title": "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}