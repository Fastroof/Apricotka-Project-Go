package entities

type OrderItem struct {
	ID           uint64  `gorm:"primaryKey;column:id" json:"-"`
	ProductID    uint64  `gorm:"column:product_id"`
	InitialPrice float64 `gorm:"column:initial_price" json:"-"`
	ExQuantity   int64   `gorm:"column:ex_quantity"`
	Price        float64 `gorm:"column:price"`
	Order        *Order  `gorm:"foreignKey:OrderID" json:"-"`
	OrderID      uint64  `gorm:"column:order_id"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
