package rule

import "github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order"

type MinValue struct {
	Value int
}

func (r MinValue) Satisfy(order order.Order) bool {
	return order.Product.Value > r.Value
}
