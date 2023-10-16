package rule

import "github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order"

type HasPaymentMethod struct {
	Method string
}

func (r HasPaymentMethod) Satisfy(order order.Order) bool {
	return order.Payment.Method == r.Method
}
