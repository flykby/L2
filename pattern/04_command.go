package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern


Команда:
	Поведенческий паттерн, который позволяет представить запрос в виде объекта.
	Из этого следует, что команда - это объект. Такие запросы, например, можно
	ставить в очередь, отменять или возобновлять. Паттерн Command отделяет объект,
	инициирующий операцию, от объекта, который знает, как ее выполнить. Единственное,
	что должен знать инициатор, это как отправить команду.

Плюсы:
	1. Клиент может работать с командой, не зная, как именно выполняется операция. Это позволяет изменять и добавлять операции, не затрагивая клиентский код.
	2. Так как команда представлена в виде объекта, вы можете легко реализовать механизмы отмены и повторения операций.
	3. Команды могут быть сохранены в истории, что позволяет воспроизводить последовательность операций или создавать снимки состояния системы.
Минусы:
	1. Увеличение числа классов.
	2. Повышенное использование памяти.
	3. Сложность реализации отмены и повторения.
*/

type Command interface {
	execute()
}

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

func NewResteraunt() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}

type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "pizzas")
}


type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	c.restaurant.CleanedDishes = c.restaurant.TotalDishes
	fmt.Println("dishes cleaned")
}

type Cook struct {
	Commands []Command
}

func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}


func main() {
	r := NewResteraunt()

	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	cooks := []*Cook {&Cook{}, &Cook{}}

	for i, task := range tasks {
		cook := cooks[i % len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}