package action

import "github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order"

type AddOrderLabel struct {
	Label string
}

func (a AddOrderLabel) Execute(order order.Order) order.Order {
	return order.AddLabel(a.Label)
}
