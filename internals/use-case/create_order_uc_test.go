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

		expectedLabels := []string{"free-shipping"}

		// ACT
		order := createOrderUC.Execute(product, payment)

		// ASSERT
		require.Equal(t, expectedLabels, order.Labels)
		require.Equal(t, product, order.Product)
		require.Equal(t, payment, order.Payment)
	})
}
