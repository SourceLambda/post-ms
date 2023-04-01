package models

type Category struct {
	ID                uint32
	Name              string `gorm:"not null"`
	Parent_CategoryID uint32 `gorm:"default:null;foreignKey:PostRefer"`
}
