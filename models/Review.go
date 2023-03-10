package models

// TableName overrides the table name and it's no more pluralized 
func (Review) TableName() string {
	return "review"
}

type Review struct {

	ID     uint32
	PostID uint32 `gorm:"not null"`
	user_name string `gorm:"not null"`
	user_email string `gorm:"not null"`
	rating uint   `gorm:"not null"`
	review_text   string `gorm:"not null"`
}
