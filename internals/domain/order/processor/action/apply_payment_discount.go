package action

import "github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order"

type ApplyPaymentDiscount struct {
	DiscountPercentage float64
}

func (a ApplyPaymentDiscount) Execute(order order.Order) order.Order {
	order.Payment = order.Payment.ApplyDiscount(a.DiscountPercentage)
	return order
}
