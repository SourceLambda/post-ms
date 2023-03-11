package models
  
// TableName overrides the table name and it's no more pluralized 
func (Post) TableName() string {
	return "post"
}

type Post struct {

	ID            uint32
	Title         string    `gorm:"not null"`
	CategoryID    uint32    `gorm:"not null"`
	Image         string    `gorm:"not null"`
	Description   string    `gorm:"not null"`
	Creation_date string `gorm:"type:date;not null"`
	Units         uint
	Price         string 	`gorm:"type:money;not null"`
	Sum_ratings   uint 		
	Num_ratings   uint 		
	Views		  uint 		
}
