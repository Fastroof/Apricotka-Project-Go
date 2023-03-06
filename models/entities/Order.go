package entities

import (
	"apricotka.com.ua/packages/models/enums"
	"time"
)

type Order struct {
	ID          uint              `gorm:"primaryKey"`
	Type        string            `gorm:"column:type" json:"-"`
	Status      enums.OrderStatus `gorm:"column:status" json:"status"`
	OrderItems  []OrderItem       `gorm:"foreignKey:OrderID"`
	Timestamp   time.Time         `gorm:"column:timestamp" json:"timestamp"`
	UserID      uint              `gorm:"column:user_id"`
	Username    string            `gorm:"column:username" json:"-"`
	Email       string            `gorm:"column:email" json:"email"`
	Phone       string            `gorm:"column:phone" json:"-"`
	Address     string            `gorm:"column:address" json:"-"`
	PaymentType string            `gorm:"column:payment_type" json:"-"`
	Payment     string            `gorm:"column:payment" json:"-"`
}

func (o *Order) TableName() string {
	return "orders"
}
