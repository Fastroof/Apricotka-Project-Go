package entities

type ProductImage struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null;unique;column:id" json:"id"`
	ProductId uint64 `gorm:"column:product_id;not null" json:"product_id"`
	File      string `gorm:"column:image_file;not null" json:"image_file"`
}

func (ProductImage) TableName() string {
	return "product_images"
}
