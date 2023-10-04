package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time" 
    )
    
    type DatasetJobtitle struct {
        ID                              uint   `json:"id" gorm:"primaryKey"` 
        Nama                            string  
        Name                            string   
        WildcardId                      int     `gorm:"column:wildcard_id"`
        WildcardEn                      int     `gorm:"column:wildcard_en"`
        ReplaceNama                     string     `gorm:"column:replace_nama"`
        ReplaceName                     string     `gorm:"column:replace_name"`
    }

	func (m *DatasetJobtitle) TableName() string {
		return "dataset_jobtitle"
	}
   
    func (b *DatasetJobtitle) GetDatasetJobtitle() ([]DatasetJobtitle, error) {
        var datasetjobtitle []DatasetJobtitle
    
        err := database.DB.  
            Select("*, REPLACE(nama, ' ', '_') AS replace_nama, REPLACE(name, ' ', '_') AS replace_name"). 
            Find(&datasetjobtitle).Error

        return datasetjobtitle, err
    }  

	func (b *DatasetJobtitle) Create(data *DatasetJobtitle) error {
		if err := database.DB.Create(data).Error; err != nil {
			return err
		}
		return nil
	}

	// UPDATE `jobtitle` SET  `jobtitle_id`='95'  WHERE name = 'Secretary'  
	// UPDATE `dataset_jobtitle` SET  `equal`='1'  WHERE id = '95' 

	// GetDatasetJobtitleByName

   
    func (b *DatasetJobtitle) GetDatasetJobtitleByNameId(name string) ([]DatasetJobtitle, error) {
        var datasetjobtitle []DatasetJobtitle
    
        err := database.DB.  
            Select("*"). 
            Where("nama = ?", name).
            Find(&datasetjobtitle).Error

        return datasetjobtitle, err
    } 
	
   
    func (b *DatasetJobtitle) GetDatasetJobtitleByNameEn(name string) ([]DatasetJobtitle, error) {
        var datasetjobtitle []DatasetJobtitle
    
        err := database.DB.  
            Select("*"). 
            Where("name = ?", name).
            Find(&datasetjobtitle).Error

        return datasetjobtitle, err
    } 
	 
    // func (j *DatasetJobtitle) UpdateDatasetJobtitle(id uint, columnName string, newValue interface{}) error {
    //     // Find the job title with the given name
    //     err := database.DB.Where("id = ?", id).First(j).Error
    //     if err != nil {
    //         return err
    //     }
    
    //     // Update the specified column with the new value
    //     err = database.DB.Model(j).Update(columnName, newValue).Error
    //     if err != nil {
    //         return err
    //     }
    
    //     return nil
    // }

	func (j *DatasetJobtitle) UpdateDatasetJobtitle(id uint, columnName string, newValue interface{}) error {
		// Find all job titles with the given ID
		var jobtitles []DatasetJobtitle
		err := database.DB.Where("id = ?", id).Find(&jobtitles).Error
		if err != nil {
			return err
		}
	
		// Update the specified column with the new value for each matching job title
		for _, jobtitle := range jobtitles {
			err := database.DB.Model(&jobtitle).Update(columnName, newValue).Error
			if err != nil {
				return err
			}
		}
	
		return nil
	}
	