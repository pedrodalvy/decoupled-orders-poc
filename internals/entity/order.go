package entity

import "github.com/pedrodalvy/decoupled-orders-poc/internals/domain/product"

const FreeShippingLabel = "free-shipping"
const FragileLabel = "fragile"
const GiftLabel = "gift"

type Order struct {
	Product product.Product
	Payment Payment
	Labels  []string
}

func (order Order) AddLabel(label string) Order {
	order.Labels = append(order.Labels, label)
	return order
}
