package useCase

import (
	"github.com/pedrodalvy/decoupled-orders-poc/internals/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateOrderUC_Execute(t *testing.T) {
	createOrderUC := NewCreateOrderUC()

	testCases := []struct {
		name     string
		product  entity.Product
		payment  entity.Payment
		expected entity.Order
	}{
		{
			name:    "should create an order with free shipping label",
			product: entity.Product{Category: "any", Value: 1001},
			payment: entity.Payment{Method: "any", Value: 1001},
			expected: entity.Order{
				Product: entity.Product{Category: "any", Value: 1001},
				Payment: entity.Payment{Method: "any", Value: 1001},
				Labels:  []string{entity.FreeShippingLabel},
			},
		},
		{
			name:    "should add fragile label when product category is appliance",
			product: entity.Product{Category: entity.ApplianceCategory, Value: 1000},
			payment: entity.Payment{Method: "any", Value: 1000},
			expected: entity.Order{
				Product: entity.Product{Category: entity.ApplianceCategory, Value: 1000},
				Payment: entity.Payment{Method: "any", Value: 1000},
				Labels:  []string{entity.FragileLabel},
			},
		},
		{
			name:    "should add gift label when product category is kids",
			product: entity.Product{Category: entity.KidsCategory, Value: 1000},
			payment: entity.Payment{Method: "any", Value: 1000},
			expected: entity.Order{
				Product: entity.Product{Category: entity.KidsCategory, Value: 1000},
				Payment: entity.Payment{Method: "any", Value: 1000},
				Labels:  []string{entity.GiftLabel},
			},
		},
		{
			name:    "should apply 10% discount when payment method is pix",
			product: entity.Product{Category: "any", Value: 1000},
			payment: entity.Payment{Method: entity.PixMethod, Value: 1000},
			expected: entity.Order{
				Product: entity.Product{Category: "any", Value: 1000},
				Payment: entity.Payment{Method: entity.PixMethod, Value: 900},
			},
		},
		{
			name:    "should apply many actions",
			product: entity.Product{Category: entity.ApplianceCategory, Value: 2000},
			payment: entity.Payment{Method: entity.PixMethod, Value: 2000},
			expected: entity.Order{
				Product: entity.Product{Category: entity.ApplianceCategory, Value: 2000},
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
