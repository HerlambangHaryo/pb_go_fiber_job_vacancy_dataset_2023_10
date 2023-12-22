package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type ApplyUrl struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *ApplyUrl) TableName() string {
		return "jobrequirement_applyurls"
	}
     
    func (b *ApplyUrl) GetApplyUrl() ([]ApplyUrl, error) {
        var applyurl []ApplyUrl
    
        err := database.DB.  
            Select("name").  
            Find(&applyurl).Error

        return applyurl, err
    }  
 