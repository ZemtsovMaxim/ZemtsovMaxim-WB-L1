package main

type Sber struct{}

func (p *Sber) PayBySber(amount float32) bool {
	// логика совершения платежа через сбер
	return true
}

type Tinkoff struct{}

func (a *Tinkoff) PayByTinkoff(amount float32) bool {
	// логика совершения платежа через тинькофф
	return true
}

// Интерфейс к которому надо адаптировать
type PaymentGateway interface {
	ProcessPayment(amount float32) bool
}

// адаптер для сбера
type SberAdapter struct {
	Sber *Sber
}

// метод для реализации интрерфейса
func (p *SberAdapter) ProcessPayment(amount float32) bool {
	return p.Sber.PayBySber(amount)
}

// адаптер для тинькофф
type TinkoffAdapter struct {
	Tinkoff *Tinkoff
}

// метод для реализации интрерфейса
func (a *TinkoffAdapter) ProcessPayment(amount float32) bool {
	return a.Tinkoff.PayByTinkoff(amount)
}

func main() {
	paymentGateway := &SberAdapter{
		Sber: &Sber{},
	}

	paymentGateway2 := &TinkoffAdapter{
		Tinkoff: &Tinkoff{},
	}

	// Sber
	paymentGateway.ProcessPayment(100)

	// Tinkoff
	paymentGateway2.ProcessPayment(100)
}

//Реализовать паттерн «адаптер» на любом примере.
