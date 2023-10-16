package rule

import "github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order"

type HasCategory struct {
	Category string
}

func (r HasCategory) Satisfy(order order.Order) bool {
	return order.Product.Category == r.Category
}
