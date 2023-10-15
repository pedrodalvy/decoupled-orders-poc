package action

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type AddFragileOrderLabel struct{}

func (a AddFragileOrderLabel) Execute(order entity.Order) entity.Order {
	return order.AddLabel("fragile")
}
