package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type ClosingDate struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *ClosingDate) TableName() string {
		return "jobrequirement_closingdates"
	}
     
    func (b *ClosingDate) GetClosingDate() ([]ClosingDate, error) {
        var closingdate []ClosingDate
    
        err := database.DB.  
            Select("name").  
            Find(&closingdate).Error

        return closingdate, err
    }  
 