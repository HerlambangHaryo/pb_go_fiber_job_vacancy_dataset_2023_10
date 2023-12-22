package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type IndustryValue struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *IndustryValue) TableName() string {
		return "jobrequirement_industryvalues"
	}
     
    func (b *IndustryValue) GetIndustryValue() ([]IndustryValue, error) {
        var industryvalue []IndustryValue
    
        err := database.DB.  
            Select("name").  
            Find(&industryvalue).Error

        return industryvalue, err
    }  
 