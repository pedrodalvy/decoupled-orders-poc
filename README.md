# Decoupled Orders POC

In this POC, the project entails developing a flexible solution for an online order problem in an e-commerce context.
The task involves applying a set of rules to each order, taking into account various order characteristics. The
existing company implementation currently comprises complex and hard-to-maintain code with numerous if/else statements.

## Requirements

### The system encompasses the following entities:

- **Product:** Category (string) and Value (int)
- **Payment:** Payment method (string) and Value (int)
- **Order:** Product, Payment, and labels (array of strings)

### The rules to be implemented are as follows:

- If the purchase is for a product with a value exceeding $ 1,000.00, add a `free-shipping` label to the order.
- If the product belongs to the `appliance` category, add a `fragile` label to the order.
- If the product belongs to the `kids` category, add a `gift` label to the order.
- If the payment method is `PIX`, provide a 10% discount.

The implemented code should be flexible so that new rules can be easily added and removed.

## Solution

To address the complex and hard-to-maintain code with numerous if/else statements, a flexible and
extensible solution was developed using a rules and actions system. 
This allows developers to easily add or remove rules without the need to modify the core codebase.

The solution consists of the following files:

```shell
internals/domain/order/processor
├── action
│    ├── add_order_label.go
│    └── apply_payment_discount.go
├── rule
│    ├── has_category.go
│    ├── has_payment_method.go
│    └── min_value.go
└── order_processor.go
```

To implement the rules, the following code is used:

```go
//...
func NewOrderProcessor() *OrderProcessor {
    orderProcessor := OrderProcessor{}
    
    orderProcessor.addConfig(rule.MinValue{Value: 1000}, action.AddOrderLabel{Label: order.FreeShippingLabel})
    orderProcessor.addConfig(rule.HasCategory{Category: product.ApplianceCategory}, action.AddOrderLabel{Label: order.FragileLabel})
    orderProcessor.addConfig(rule.HasCategory{Category: product.KidsCategory}, action.AddOrderLabel{Label: order.GiftLabel})
    orderProcessor.addConfig(rule.HasPaymentMethod{Method: payment.PixMethod}, action.ApplyPaymentDiscount{DiscountPercentage: 10})
    
    return &orderProcessor
}
//...
```
