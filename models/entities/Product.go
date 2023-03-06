package entities

type Product struct {
	ID                     uint64                  `gorm:"primaryKey;column:id"`
	CategoryID             uint64                  `gorm:"column:category_id"`
	IncomingProductDetails []IncomingProductDetail `gorm:"foreignKey:ProductID" json:"-"`
	Name                   string                  `gorm:"column:name;not null;unique"`
	Price                  float64                 `gorm:"column:price;not null"`
	Info                   string                  `gorm:"column:info;type:text"`
}

func (Product) TableName() string {
	return "products"
}
