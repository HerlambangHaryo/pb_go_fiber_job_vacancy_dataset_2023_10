package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type PostedDate struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *PostedDate) TableName() string {
		return "posted_dates"
	}
     
    func (b *PostedDate) GetPostedDate() ([]PostedDate, error) {
        var posteddate []PostedDate
    
        err := database.DB.  
            Select("name").  
            Find(&posteddate).Error

        return posteddate, err
    }  
 