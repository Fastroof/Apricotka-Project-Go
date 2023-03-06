package entities

type Gallery struct {
	ID        uint64 `gorm:"primaryKey;column:group_id"`
	GroupName string `gorm:"column:group_name"`
}

func (Gallery) TableName() string {
	return "gallery"
}
