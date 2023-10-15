package rule

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type MinValue struct {
	Value int
}

func (r MinValue) Satisfy(order entity.Order) bool {
	return order.Product.Value > r.Value
}
