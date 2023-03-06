package enums

type OrderStatus string

const (
	Rejected    OrderStatus = "rejected"
	Pending     OrderStatus = "pending"
	Initialized OrderStatus = "initialized"
	Sent        OrderStatus = "sent"
	Completed   OrderStatus = "completed"
)
