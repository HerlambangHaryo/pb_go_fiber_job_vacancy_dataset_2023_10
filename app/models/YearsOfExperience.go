package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type YearsOfExperience struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *YearsOfExperience) TableName() string {
		return "years_of_experiences"
	}
     
    func (b *YearsOfExperience) GetYearsOfExperience() ([]YearsOfExperience, error) {
        var yearsofexperience []YearsOfExperience
    
        err := database.DB.  
            Select("name").  
            Find(&yearsofexperience).Error

        return yearsofexperience, err
    }  
 