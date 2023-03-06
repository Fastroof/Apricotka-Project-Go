package entities

type GalleryImage struct {
	ID        uint64  `gorm:"primaryKey;column:id" json:"id"`
	Gallery   Gallery `gorm:"foreignKey:GalleryId" json:"-"`
	GalleryId uint64  `gorm:"column:gallery_id" json:"gallery_id"`
	File      string  `gorm:"column:file" json:"file"`
}

func (GalleryImage) TableName() string {
	return "gallery_images"
}
