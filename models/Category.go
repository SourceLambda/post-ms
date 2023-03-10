package models
 
// TableName overrides the table name and it's no more pluralized 
func (Category) TableName() string {
	return "category"
}

type Category struct {
	ID    uint32
	Name  string `gorm:"not null"`
	Parent_CategoryID uint32 `gorm:"default:null;foreignKey:PostRefer"`
}
