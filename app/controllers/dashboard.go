// book.go
package controllers

    import (
        "github.com/gofiber/fiber/v2" 
    ) 
  
    func GetDashboard(c *fiber.Ctx) error { 
        // --------------------------------------------------------------  
        // --------------------------------------------------------------    
        // --------------------------------------------------------------
            return c.Render("contents/dashboard/GetDashboard", fiber.Map{
                "Title": "League Not Started", 
				"Content": "Dashboard",
            },"templates/studiov30/pageblank")
        // --------------------------------------------------------------
    }