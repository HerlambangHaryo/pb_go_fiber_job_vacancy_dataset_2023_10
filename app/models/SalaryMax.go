package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type SalaryMax struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *SalaryMax) TableName() string {
		return "salary_max"
	}
     
    func (b *SalaryMax) GetSalaryMax() ([]SalaryMax, error) {
        var salarymax []SalaryMax
    
        err := database.DB.  
            Select("name").  
            Find(&salarymax).Error

        return salarymax, err
    }  
 