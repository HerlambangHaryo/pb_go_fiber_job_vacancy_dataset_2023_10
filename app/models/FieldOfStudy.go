package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type FieldOfStudy struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *FieldOfStudy) TableName() string {
		return "field_of_studies"
	}
     
    func (b *FieldOfStudy) GetFieldOfStudy() ([]FieldOfStudy, error) {
        var fieldofstudy []FieldOfStudy
    
        err := database.DB.  
            Select("name").  
            Find(&fieldofstudy).Error

        return fieldofstudy, err
    }  
 