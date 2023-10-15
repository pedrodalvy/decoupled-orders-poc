package rule

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type HasCategory struct {
	Category string
}

func (r HasCategory) Satisfy(order entity.Order) bool {
	return order.Product.Category == r.Category
}
