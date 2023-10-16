package useCase

import (
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/order/processor"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/payment"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/product"
)

type CreateOrderUC struct {
	orderProcessor *processor.OrderProcessor
}

func NewCreateOrderUC() *CreateOrderUC {
	return &CreateOrderUC{
		orderProcessor: processor.NewOrderProcessor(),
	}
}

func (c CreateOrderUC) Execute(product product.Product, payment payment.Payment) order.Order {
	o := order.Order{Product: product, Payment: payment}

	for _, config := range c.orderProcessor.Config {
		if config.Rule.Satisfy(o) {
			o = config.Action.Execute(o)
		}
	}

	return o
}
