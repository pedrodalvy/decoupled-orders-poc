package processor

import (
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order/processor/action"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order/processor/rule"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/payment"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/product"
)

type Rule interface {
	Satisfy(order order.Order) bool
}

type Action interface {
	Execute(order order.Order) order.Order
}

type OrderProcessor struct {
	Config []struct {
		Rule   Rule
		Action Action
	}
}

func NewOrderProcessor() *OrderProcessor {
	orderProcessor := OrderProcessor{}

	orderProcessor.addConfig(rule.MinValue{Value: 1000}, action.AddOrderLabel{Label: order.FreeShippingLabel})
	orderProcessor.addConfig(rule.HasCategory{Category: product.ApplianceCategory}, action.AddOrderLabel{Label: order.FragileLabel})
	orderProcessor.addConfig(rule.HasCategory{Category: product.KidsCategory}, action.AddOrderLabel{Label: order.GiftLabel})
	orderProcessor.addConfig(rule.HasPaymentMethod{Method: payment.PixMethod}, action.ApplyPaymentDiscount{DiscountPercentage: 10})

	return &orderProcessor
}

func (p *OrderProcessor) addConfig(rule Rule, action Action) {
	p.Config = append(p.Config, struct {
		Rule   Rule
		Action Action
	}{rule, action})
}
