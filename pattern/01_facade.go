package pattern



/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад:
	Позволяет скрыть сложность системы путём сведения всех возможных внешних вызовов к одному объекту,
	делегирующему их соответствующим объектам системы.

Плюсы:
	Изолирует клиентов от поведения сложной системы. Сам интерфейс фасада очень прост,
	тем самым мы скрываем сложную структуру поведения от пользователя
Минусы:
	Сам интерфейс фасада может стать супер-классом. Это значит, что обьект может быть
	использован во всей системе. И вся система будет привязана к нему, все последующие
	функции будут проходить через него

Применяется:
	1) Облегченное взаимодействие с монолитом
	2) Упрощение работы со сложной системой
	3) Минимизация связанности

В данном коде фасадом является метод Shop.Sell, который вызывает в себе сложную системы из нескольких
вызовов методов других структур. Но при используемые в нем функции остаются доступными для "опытных" юзеров.

Реализуем поведение покупки товара:
*/

import "fmt"

// Структура товара в магазине
type Product struct {
	Name  string
	Price float64
}

// Структура магазина
type Shop struct {
	Address string
	Product []Product
}

func (shop *Shop) Sell(user User, productName string) bool {
	
	if !user.Card.CheckBalance() {
		return false
	}

	for _, p := range shop.Product {
		if p.Name != productName {
			continue
		}
		if p.Price > user.GetBalance() {
			return false
		}
		if p.Price <= user.GetBalance() {
			user.Card.Balance -= p.Price
			return true
		}
	}
	return false
}

type Bank struct {
	Name  string
	Cards []Card
}

func (bank *Bank) CheckBalance(number string) bool {
	for _, card := range bank.Cards {
		if card.Number != number {
			continue
		}
		if card.Balance <= 0 {
			return false
		}
	}
	return true
}

type Card struct {
	Number  string
	Balance float64
	Bank    *Bank
}

func (card *Card) CheckBalance() bool {
	return card.Bank.CheckBalance(card.Number)
}

type User struct {
	Name string
	Card *Card
}

func (user *User) GetBalance() float64 {
	return user.Card.Balance
}


func main() {
	bank := &Bank{Name: "Бета-Банк"}

	product1 := &Product{
		Name: "product1",
		Price: 5001,
	}
	product2 := &Product{
		Name: "product2",
		Price: 4000,
	}

	shop := &Shop{
		Address: "Lenina street 42",

	}
	shop.Product = append(shop.Product, *product1, *product2)

	card1 := &Card{
		Number: "4024442120",
		Balance: 10000,
		Bank: bank,
	}
	card2 := &Card{
		Number: "4024441890",
		Balance: 5000,
		Bank: bank,
	}
	bank.Cards = append(bank.Cards, *card1, *card2)

	user1 := &User{
		Name: "user1",
		Card: card1,
	}
	user2 := &User{
		Name: "user2",
		Card: card2,
	}

	fmt.Println(shop.Sell(*user1, product1.Name))
	fmt.Println(shop.Sell(*user2, product1.Name))


}