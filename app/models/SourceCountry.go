package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type SourceCountry struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *SourceCountry) TableName() string {
		return "sourcecountries"
	}
     
    func (b *SourceCountry) GetSourceCountry() ([]SourceCountry, error) {
        var sourcecountry []SourceCountry
    
        err := database.DB.  
            Select("name").  
            Find(&sourcecountry).Error

        return sourcecountry, err
    }  
 