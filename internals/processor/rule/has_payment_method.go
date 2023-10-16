package rule

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type HasPaymentMethod struct {
	Method string
}

func (r HasPaymentMethod) Satisfy(order entity.Order) bool {
	return order.Payment.Method == r.Method
}
