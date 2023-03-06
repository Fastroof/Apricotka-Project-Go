package entities

import (
	"time"
)

type IncomingProductDetail struct {
	ID           uint64    `gorm:"primaryKey;column:id"`
	Product      Product   `gorm:"foreignKey:ProductID" json:"-"`
	ProductID    uint64    `gorm:"column:product_id"`
	Supplier     string    `gorm:"column:supplier"`
	InitialPrice float64   `gorm:"column:initial_price;not null"`
	Quantity     int64     `gorm:"column:quantity;not null"`
	Timestamp    time.Time `gorm:"column:timestamp;not null"`
}

func (IncomingProductDetail) TableName() string {
	return "incoming_product_details"
}
