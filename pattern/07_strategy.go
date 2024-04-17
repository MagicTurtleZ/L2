package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
*/

/*
Разработка приложения для обработки платежей, где разные стратегии могут быть применены в зависимости
от типа платежа (например, кредитная карта, электронные деньги или банковский перевод).
*/

// стратегия определяет интерфейс, общий для всех вариаций алгоритма
type PaymentStrategy interface {
	payment(amout float64)
}

// конкретные стратегии реализуют различные вариации алгоритма.
type CreditCard struct {
	cardNumber     string
	expirationDate string
	cvc            string
}

func (cc *CreditCard) payment(amout float64) {
	fmt.Println("implementation of a credit card payment system")
}

type EPSystem struct {
	email    string
	password string
}

func (eps *EPSystem) payment(amout float64) {
	fmt.Println("implementation of electronic payment systems")
}

type BankTransfer struct {
	accountNumber string
	routingNumber string
}

func (bt *BankTransfer) payment(amout float64) {
	fmt.Println("implementation of a payment system by bank transfer")
}

// контекст хранит ссылку на объект конкретной стратегии, работая с ним через общий интерфейс стратегий.
type Buyer struct {
	paymentStrategy PaymentStrategy
}

func (b *Buyer) SetStrategy(ps PaymentStrategy) {
	b.paymentStrategy = ps
}

func (b *Buyer) Pay(amout float64) {
	b.paymentStrategy.payment(amout)
}

// клиентский код
func ExampleStrategy() {
	cc := new(CreditCard)

	buyer := &Buyer{paymentStrategy: cc}
	buyer.Pay(123)

	eps := new(EPSystem)

	buyer.SetStrategy(eps)
	buyer.Pay(456)

	bt := new(BankTransfer)

	buyer.SetStrategy(bt)
	buyer.Pay(789)
}

/*
	Плюсы:
1. горячая замена алгоритмов на лету
2. изолирует код и данные алгоритмов от остальных классов
	Минусы:
1. усложняет программу за счёт дополнительных классов
2. клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую
*/