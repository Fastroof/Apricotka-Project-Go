package entities

type Gallery struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement;not null;unique;column:id" json:"id"`
	Title string `gorm:"column:title;not null" json:"title"`
}

func (Gallery) TableName() string {
	return "gallery"
}
