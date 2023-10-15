package processor

import (
	"github.com/pedrodalvy/decoupled-orders-poc/internals/entity"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/processor/action"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/processor/rule"
)

type Rule interface {
	Satisfy(order entity.Order) bool
}

type Action interface {
	Execute(order entity.Order) entity.Order
}

type OrderProcessor struct {
	Config []struct {
		Rule   Rule
		Action Action
	}
}

func NewOrderProcessor() *OrderProcessor {
	orderProcessor := OrderProcessor{}

	orderProcessor.addConfig(rule.MinValue{Value: 1000}, action.AddFreeShippingOrderLabel{})

	return &orderProcessor
}

func (p *OrderProcessor) addConfig(rule Rule, action Action) {
	p.Config = append(p.Config, struct {
		Rule   Rule
		Action Action
	}{rule, action})
}