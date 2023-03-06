package entities

type GalleryImage struct {
	ID      uint64 `gorm:"primaryKey;column:image_id"`
	GroupId uint64 `gorm:"column:group_id"`
	File    string `gorm:"column:image_file"`
}

func (GalleryImage) TableName() string {
	return "gallery_images"
}
