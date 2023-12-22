package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type CareerLevel struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *CareerLevel) TableName() string {
		return "career_levels"
	}
     
    func (b *CareerLevel) GetCareerLevel() ([]CareerLevel, error) {
        var careerlevel []CareerLevel
    
        err := database.DB.  
            Select("name").  
            Find(&careerlevel).Error

        return careerlevel, err
    }  
 