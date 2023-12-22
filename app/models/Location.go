package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type Location struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *Location) TableName() string {
		return "locations"
	}
     
    func (b *Location) GetLocation() ([]Location, error) {
        var location []Location
    
        err := database.DB.  
            Select("name").  
            Find(&location).Error

        return location, err
    }  
 