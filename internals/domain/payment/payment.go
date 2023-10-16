package payment

const PixMethod = "pix"

type Payment struct {
	Method string
	Value  int
}

func (payment Payment) ApplyDiscount(percentage float64) Payment {
	payment.Value = int(float64(payment.Value) * (1 - (percentage / 100)))
	return payment
}
