package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type SalaryMin struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *SalaryMin) TableName() string {
		return "salary_mins"
	}
     
    func (b *SalaryMin) GetSalaryMin() ([]SalaryMin, error) {
        var salarymin []SalaryMin
    
        err := database.DB.  
            Select("name").  
            Find(&salarymin).Error

        return salarymin, err
    }  
 