package action

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type AddOrderLabel struct {
	Label string
}

func (a AddOrderLabel) Execute(order entity.Order) entity.Order {
	return order.AddLabel(a.Label)
}
