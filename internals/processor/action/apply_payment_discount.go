package action

import "github.com/pedrodalvy/decoupled-orders-poc/internals/entity"

type ApplyPaymentDiscount struct {
	DiscountPercentage float64
}

func (a ApplyPaymentDiscount) Execute(order entity.Order) entity.Order {
	order.Payment = order.Payment.ApplyDiscount(a.DiscountPercentage)
	return order
}
