package useCase

import (
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/product"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/entity"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/processor"
)

type CreateOrderUC struct {
	orderProcessor *processor.OrderProcessor
}

func NewCreateOrderUC() *CreateOrderUC {
	return &CreateOrderUC{
		orderProcessor: processor.NewOrderProcessor(),
	}
}

func (c CreateOrderUC) Execute(product product.Product, payment entity.Payment) entity.Order {
	order := entity.Order{Product: product, Payment: payment}

	for _, config := range c.orderProcessor.Config {
		if config.Rule.Satisfy(order) {
			order = config.Action.Execute(order)
		}
	}

	return order
}
