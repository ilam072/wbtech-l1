package main

import "fmt"

// Адаптер — это структурный паттерн проектирования, который
// позволяет объектам с несовместимыми интерфейсами работать вместе.

// +:
// 1. Отделяет и скрывает от клиента подробности преобразования различных интерфейсов
// -:
// 1. Усложняет код программы из-за введения дополнительных классов

// Старый интерфейс
type PaymentProcessor interface {
	Pay(amount float64) error
}

// Старая реализация
type OldPaymentService struct{}

func (o *OldPaymentService) Pay(amount float64) error {
	fmt.Printf("Оплата через старый сервис: %.2f\n", amount)
	return nil
}

// Новый сервис с другим методом
type NewPaymentService struct{}

func (n *NewPaymentService) SendPayment(currency string, value float64) error {
	fmt.Printf("Оплата через новый сервис: %.2f %s\n", value, currency)
	return nil
}

// Адаптер
type PaymentAdapter struct {
	s *NewPaymentService
}

func (p *PaymentAdapter) Pay(amount float64) error {
	// Преобразуем вызов под новый формат
	return p.s.SendPayment("USD", amount)
}

func main() {
	// Старый способ
	var processor PaymentProcessor = &OldPaymentService{}
	processor.Pay(100)

	// Новый способ через адаптер
	newService := &NewPaymentService{}
	processor = &PaymentAdapter{s: newService}
	processor.Pay(250)
}
