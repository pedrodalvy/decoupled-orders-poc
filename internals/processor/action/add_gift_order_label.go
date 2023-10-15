package action

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type AddGiftOrderLabel struct{}

func (a AddGiftOrderLabel) Execute(order entity.Order) entity.Order {
	return order.AddLabel("gift")
}
