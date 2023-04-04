package models

type Review struct {
	ID          uint32
	PostID      uint32 `gorm:"not null"`
	User_name   string `gorm:"not null"`
	User_email  string `gorm:"not null"`
	Rating      uint   `gorm:"not null"`
	Review_text string `gorm:"not null"`
}

type OldReview struct {
	ID          uint32
	PostID      uint32 `gorm:"not null"`
	User_name   string `gorm:"not null"`
	User_email  string `gorm:"not null"`
	Rating      uint   `gorm:"not null"`
	Review_text string `gorm:"not null"`
	OldRating   uint   `gorm:"not null"`
}

type RequestBody struct {
	OldRating uint
	PostID    uint32
}
