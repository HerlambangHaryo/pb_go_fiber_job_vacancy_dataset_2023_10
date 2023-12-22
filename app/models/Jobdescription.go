package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type Jobdescription struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *Jobdescription) TableName() string {
		return "jobdescription"
	}
     
    func (b *Jobdescription) GetJobdescription() ([]Jobdescription, error) {
        var jobdescription []Jobdescription
    
        err := database.DB.  
            Select("name").  
            Find(&jobdescription).Error

        return jobdescription, err
    }  
	
	// func (j *DatasetJobtitle) GetJobdescriptionSetDone(id uint, columnName string, newValue interface{}) error {
	// 	// Find all job titles with the given ID
	// 	var jobtitles []DatasetJobtitle
	// 	err := database.DB.Where("id = ?", id).Find(&jobtitles).Error
	// 	if err != nil {
	// 		return err
	// 	}
	
	// 	// Update the specified column with the new value for each matching job title
	// 	for _, jobtitle := range jobtitles {
	// 		err := database.DB.Model(&jobtitle).Update(columnName, newValue).Error
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	
	// 	return nil
	// }

	
	func (j *Jobdescription) GetJobdescriptionSetDone(id string, columnName string, newValue interface{}) error {
		err := database.DB.
			Model(&j).
			Where("id = ?", id).
			Update(columnName, newValue).
			Error
		return err
	}
	
 