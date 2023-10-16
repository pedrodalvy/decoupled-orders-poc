package useCase

import (
	"github.com/pedrodalvy/decoupled-orders-poc/internals/entity"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateOrderUC_Execute(t *testing.T) {
	createOrderUC := NewCreateOrderUC()

	t.Run("should create an order with free shipping label", func(t *testing.T) {
		// ARRANGE
		product := entity.Product{Category: "any", Value: 1001}
		payment := entity.Payment{Method: "any", Value: 1001}

		expectedLabels := []string{entity.FreeShippingLabel}

		// ACT
		order := createOrderUC.Execute(product, payment)

		// ASSERT
		require.Equal(t, expectedLabels, order.Labels)
		require.Equal(t, product, order.Product)
		require.Equal(t, payment, order.Payment)
	})

	t.Run("should add fragile label when product category is appliance", func(t *testing.T) {
		// ARRANGE
		product := entity.Product{Category: entity.ApplianceCategory, Value: 1000}
		payment := entity.Payment{Method: "any", Value: 1000}

		expectedLabels := []string{entity.FragileLabel}

		// ACT
		order := createOrderUC.Execute(product, payment)

		// ASSERT
		require.Equal(t, expectedLabels, order.Labels)
		require.Equal(t, product, order.Product)
		require.Equal(t, payment, order.Payment)
	})

	t.Run("should add gift label when product category is kids", func(t *testing.T) {
		// ARRANGE
		product := entity.Product{Category: entity.KidsCategory, Value: 1000}
		payment := entity.Payment{Method: "any", Value: 1000}

		expectedLabels := []string{entity.GiftLabel}

		// ACT
		order := createOrderUC.Execute(product, payment)

		// ASSERT
		require.Equal(t, expectedLabels, order.Labels)
		require.Equal(t, product, order.Product)
		require.Equal(t, payment, order.Payment)
	})

	t.Run("should add 10% discount when payment method is pix", func(t *testing.T) {
		// ARRANGE
		product := entity.Product{Category: "any", Value: 1000}
		payment := entity.Payment{Method: entity.PixMethod, Value: 1000}

		// ACT
		order := createOrderUC.Execute(product, payment)

		// ASSERT
		require.Empty(t, order.Labels)
		require.Equal(t, product, order.Product)
		require.Equal(t, payment.Method, order.Payment.Method)
		require.Equal(t, 900, order.Payment.Value)
	})
}
