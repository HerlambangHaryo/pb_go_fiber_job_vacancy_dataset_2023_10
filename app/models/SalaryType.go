package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type SalaryType struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *SalaryType) TableName() string {
		return "salary_types"
	}
     
    func (b *SalaryType) GetSalaryType() ([]SalaryType, error) {
        var salarytype []SalaryType
    
        err := database.DB.  
            Select("name").  
            Find(&salarytype).Error

        return salarytype, err
    }  
 