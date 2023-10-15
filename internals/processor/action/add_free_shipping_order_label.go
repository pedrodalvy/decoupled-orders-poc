package action

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type AddFreeShippingOrderLabel struct{}

func (a AddFreeShippingOrderLabel) Execute(order entity.Order) entity.Order {
	return order.AddLabel("free-shipping")
}
