package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type IndustryRequirement struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *IndustryRequirement) TableName() string {
		return "industry_requirements"
	}
     
    func (b *IndustryRequirement) GetIndustryRequirement() ([]IndustryRequirement, error) {
        var industryrequirement []IndustryRequirement
    
        err := database.DB.  
            Select("name").  
            Find(&industryrequirement).Error

        return industryrequirement, err
    }  
 