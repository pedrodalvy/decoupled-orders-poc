package useCase

import (
	"github.com/pedrodalvy/decoupled-orders-poc/internals/domain/product"
	"github.com/pedrodalvy/decoupled-orders-poc/internals/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateOrderUC_Execute(t *testing.T) {
	createOrderUC := NewCreateOrderUC()

	testCases := []struct {
		name     string
		product  product.Product
		payment  entity.Payment
		expected entity.Order
	}{
		{
			name:    "should create an order with free shipping label",
			product: product.Product{Category: "any", Value: 1001},
			payment: entity.Payment{Method: "any", Value: 1001},
			expected: entity.Order{
				Product: product.Product{Category: "any", Value: 1001},
				Payment: entity.Payment{Method: "any", Value: 1001},
				Labels:  []string{entity.FreeShippingLabel},
			},
		},
		{
			name:    "should add fragile label when product category is appliance",
			product: product.Product{Category: product.ApplianceCategory, Value: 1000},
			payment: entity.Payment{Method: "any", Value: 1000},
			expected: entity.Order{
				Product: product.Product{Category: product.ApplianceCategory, Value: 1000},
				Payment: entity.Payment{Method: "any", Value: 1000},
				Labels:  []string{entity.FragileLabel},
			},
		},
		{
			name:    "should add gift label when product category is kids",
			product: product.Product{Category: product.KidsCategory, Value: 1000},
			payment: entity.Payment{Method: "any", Value: 1000},
			expected: entity.Order{
				Product: product.Product{Category: product.KidsCategory, Value: 1000},
				Payment: entity.Payment{Method: "any", Value: 1000},
				Labels:  []string{entity.GiftLabel},
			},
		},
		{
			name:    "should apply 10% discount when payment method is pix",
			product: product.Product{Category: "any", Value: 1000},
			payment: entity.Payment{Method: entity.PixMethod, Value: 1000},
			expected: entity.Order{
				Product: product.Product{Category: "any", Value: 1000},
				Payment: entity.Payment{Method: entity.PixMethod, Value: 900},
			},
		},
		{
			name:    "should apply many actions",
			product: product.Product{Category: product.ApplianceCategory, Value: 2000},
			payment: entity.Payment{Method: entity.PixMethod, Value: 2000},
			expected: entity.Order{
				Product: product.Product{Category: product.ApplianceCategory, Value: 2000},
				Payment: entity.Payment{Method: entity.PixMethod, Value: 1800},
				Labels:  []string{entity.FreeShippingLabel, entity.FragileLabel},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// ACT
			order := createOrderUC.Execute(testCase.product, testCase.payment)

			// ASSERT
			require.Equal(t, testCase.expected, order)
		})
	}
}
