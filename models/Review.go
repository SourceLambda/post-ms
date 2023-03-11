package models

// TableName overrides the table name and it's no more pluralized 
func (Review) TableName() string {
	return "review"
}

type Review struct {

	ID     uint32
	PostID uint32 `gorm:"not null"`
	User_name string `gorm:"not null"`
	User_email string `gorm:"not null"`
	Rating uint   `gorm:"not null"`
	Review_text   string `gorm:"not null"`
}
