package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time" 
    )
    
    type DatasetCareerLevel struct {
        ID                              uint   `json:"id" gorm:"primaryKey"` 
        Nama                            string   
    }

	func (m *DatasetCareerLevel) TableName() string {
		return "dataset_career_levels"
	}
   
    func (b *DatasetCareerLevel) GetDatasetCareerLevel() ([]DatasetCareerLevel, error) {
        var datasetcareerlevel []DatasetCareerLevel
    
        err := database.DB.  
            Select("*"). 
            Find(&datasetcareerlevel).Error

        return datasetcareerlevel, err
    }  