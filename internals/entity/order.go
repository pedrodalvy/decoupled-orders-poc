package entity

const FreeShippingLabel = "free-shipping"
const FragileLabel = "fragile"
const GiftLabel = "gift"

type Order struct {
	Product Product
	Payment Payment
	Labels  []string
}

func (order Order) AddLabel(label string) Order {
	order.Labels = append(order.Labels, label)
	return order
}
