package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Стратегия:
	Поведенческий паттерн проектирования, позволяющий выбор поведения алгоритма в ходе
	исполнения. Этот паттерн определяет алгоритмы, инкапсулирует их и использует их
	взаимозаменяемо. Паттерн Strategy позволяет подменять алгоритмы без участия клиентов,
	которые используют эти алгоритмы.

Плюсы:
	1) Динамическое определение, какой алгоритм будет запущен.
	2) Инкапсуляция. Отделен код конкретной стратегии от остального кода.

Минусы:
	1) Увеличение размера кода
	2) Требует понимание об

Применяется:
	1) Когда имеется много родственных классов, отличающихся только поведением;
	2) Когда нужно иметь несколько разных вариантов алгоритма;
	3) Когда в классе определено много поведений, что представлено разветвленными условными операторами.

В данном коде паттерн страгея реализуется для разделения бизнес-логики обработки заказа и оплаты заказа.
В метод processOrder нам передается обьект payment, о котором он не знает ничего, кроме того, что у него
есть метод Pay()

Реализуем варианты способов оплаты товара с помощью паттерна стратегия:
*/

type Payment interface {
	Pay() error
}

type cardPayment struct {
	number         string
	validityPeriod string
	cvv            string
}

func NewCardPayment(number, validityPeriod, cvv string) Payment {
	return &cardPayment{
		number:         number,
		validityPeriod: validityPeriod,
		cvv:            cvv,
	}
}

// Реализация оплаты с помощью банковской карты
func (p *cardPayment) Pay() error {

	return nil
}

type sbpPayment struct {
}

func NewSbpPayment(phoneNumber string) Payment {
	return &sbpPayment{}
}

// Реализация оплаты с помощью СБП
func (p *sbpPayment) Pay() error {

	return nil
}

type walletPayment struct {
}

func NewWalletPayment(userId string) Payment {
	return &walletPayment{}
}

// Реализация оплаты с помощью кошелька на сайте
func (p *walletPayment) Pay() error {

	return nil
}

func processOrder(payment Payment) {
	err := payment.Pay()
	if err != nil {
		fmt.Println("Ошибка оплаты!")
		return
	}
	fmt.Println("Оплата успешна!")
}

func main() {
	var payment Payment
	payWay := 2

	switch payWay {
	case 1:
		payment = NewCardPayment("4300 1234 4432 9906", "12/31", "931")
	case 2:
		payment = NewSbpPayment("+7(943) 129 55-75")
	case 3:
		payment = NewWalletPayment("4d2a-12b4-b79d-dbbd")
	}

	processOrder(payment)

}
