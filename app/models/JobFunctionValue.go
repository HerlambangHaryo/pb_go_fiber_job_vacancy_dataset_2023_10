package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time"  
    )
    
    type JobFunctionValue struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"` 
    }

	func (m *JobFunctionValue) TableName() string {
		return "jobfunctionvalues"
	}
     
    func (b *JobFunctionValue) GetJobFunctionValue() ([]JobFunctionValue, error) {
        var jobfunctionvalue []JobFunctionValue
    
        err := database.DB.  
            Select("name").  
            Find(&jobfunctionvalue).Error

        return jobfunctionvalue, err
    }  
 