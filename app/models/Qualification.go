package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type Qualification struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *Qualification) TableName() string {
		return "career_levels"
	}
     
    func (b *Qualification) GetQualification() ([]Qualification, error) {
        var qualification []Qualification
    
        err := database.DB.  
            Select("name").  
            Find(&qualification).Error

        return qualification, err
    }  
 