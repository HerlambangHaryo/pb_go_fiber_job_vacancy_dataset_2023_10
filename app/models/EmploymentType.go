package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type EmploymentType struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *EmploymentType) TableName() string {
		return "employment_types"
	}
     
    func (b *EmploymentType) GetEmploymentType() ([]EmploymentType, error) {
        var employmenttype []EmploymentType
    
        err := database.DB.  
            Select("name").  
            Find(&employmenttype).Error

        return employmenttype, err
    }  
 