package models

type Post struct {
	ID            uint32
	Title         string  `gorm:"not null"`
	CategoryID    uint32  `gorm:"not null"`
	Image         string  `gorm:"not null"`
	Description   string  `gorm:"not null"`
	Creation_date string  `gorm:"type:date;not null"`
	Units         uint    `gorm:"not null"`
	Price         float32 `gorm:"type:money;not null"`
	Sum_ratings   uint
	Num_ratings   uint
	Views         uint
}
