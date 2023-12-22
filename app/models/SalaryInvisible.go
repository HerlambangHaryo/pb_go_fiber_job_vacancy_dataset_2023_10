package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type SalaryInvisible struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *SalaryInvisible) TableName() string {
		return "salary_isvisibles"
	}
     
    func (b *SalaryInvisible) GetSalaryInvisible() ([]SalaryInvisible, error) {
        var salaryinvisible []SalaryInvisible
    
        err := database.DB.  
            Select("name").  
            Find(&salaryinvisible).Error

        return salaryinvisible, err
    }  
 