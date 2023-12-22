package models

    import ( 
        "github.com/HerlambangHaryo/pb_go_fiber_job_vacancy_dataset_2023_10/platform/database" 
        // "time" 
        "gorm.io/gorm"
    )
    
    type Jobtitle struct {
        ID                              uint    `json:"id" gorm:"primaryKey"` 
        Name                            string  `gorm:"column:name"`
        Counter                         int   
		Percentage                      string  `gorm:"column:percentagex"`
		TotalRows                       string  `gorm:"column:total_rows"` 
        ReplaceName                     string  `gorm:"column:replace_name"`
    }

	func (m *Jobtitle) TableName() string {
		return "jobtitle"
	}
     
    func (b *Jobtitle) GetJobtitle() ([]Jobtitle, error) {
        var jobtitle []Jobtitle
    
        err := database.DB.  
            Select("name, COUNT(name) AS counter, REPLACE(name, ' ', '_') AS replace_name").
            Where("jobtitle_id IS NULL").
            Group("name").
            Having("COUNT(name) > 1").
            Order("counter DESC").
            Limit(100).
            Find(&jobtitle).Error

        return jobtitle, err
    }  

    func (b *Jobtitle) GetJobtitleWildcard(nama string) ([]Jobtitle, error) {
        var jobtitle []Jobtitle
    
        err := database.DB.   
            Select("name, REPLACE(name, ' ', '_') AS replace_name").
            Where("name LIKE ?", "%"+nama+"%"). 
            Where("jobtitle_id IS NULL").
            Find(&jobtitle).Error

        return jobtitle, err
    }   
 
    func (b *Jobtitle) GetJobtitleTotalRows() ([]Jobtitle, error) {
        var jobtitle []Jobtitle
    
        err := database.DB.  
            Select("COUNT(*) AS total_rows"). 
            Find(&jobtitle).Error

        return jobtitle, err
    }  

    func (b *Jobtitle) GetJobtitlePercentage() ([]Jobtitle, error) {
        var jobtitle []Jobtitle
    
        err := database.DB.  
            Select("(SELECT COUNT(*) from jobtitle where jobtitle_id is not null) / COUNT(*) * 100 AS percentagex"). 
            Find(&jobtitle).Error

        return jobtitle, err
    }   

    func (j *Jobtitle) UpdateJobtitlex(newJobtitleID uint, name string, db *gorm.DB) error {
        return db.Model(&Jobtitle{}).
            Where("name = ?", name).
            Update("jobtitle_id", newJobtitleID).Error
    } 

    func (j *Jobtitle) UpdateJobtitle(name string, columnName string, newValue interface{}) error {
        // Find all job titles with the given name
        var jobtitles []Jobtitle
        err := database.DB.Where("name = ?", name).Find(&jobtitles).Error
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