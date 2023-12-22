package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type SalaryCurrency struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *SalaryCurrency) TableName() string {
		return "salary_currencies"
	}
     
    func (b *SalaryCurrency) GetSalaryCurrency() ([]SalaryCurrency, error) {
        var salarycurrency []SalaryCurrency
    
        err := database.DB.  
            Select("name").  
            Find(&salarycurrency).Error

        return salarycurrency, err
    }  
 