package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Фабричный метод:
	Пораждающий шаблон проектирования, который определяет общий интерфейс создания
	объектов в родительском классе и позволяющий изменять создаваемые объекты в
	дочерних классах.

Плюсы:


Минусы:


*/

type iTransport interface {
	// Set name of transport
	setName(n string)
	// Get name of transport
	getName() string
	// Set speed of transport
	setSpeed(s uint)
	// Get speed of transport
	getSpeed() uint
}

type transport struct {
	name  string
	speed uint
}

// implement interface
func (t *transport) setName(n string) {
	t.name = n
}

func (t *transport) getName() string {
	return t.name
}

func (t *transport) setSpeed(s uint) {
	t.speed = s
}

func (t *transport) getSpeed() uint {
	return t.speed
}

type electricScooter struct {
	transport
}

func newElectricScooter() iTransport {
	return &electricScooter{
		transport: transport{
			name:  "Scooter",
			speed: 4,
		},
	}
}

type quadcopter struct {
	transport
}

func newQuadcopter() iTransport {
	return &quadcopter{
		transport: transport{
			name:  "Quadcopter",
			speed: 14,
		},
	}
}

func getTransport(tt string) (iTransport, error) {
	if tt == "scooter" {
		return newElectricScooter(), nil
	}
	if tt == "quadcopter" {
		return newQuadcopter(), nil
	}
	return nil, fmt.Errorf("Wrong type")
}

func main() {

	scooter, _ := getTransport("scooter")
	quad, _ := getTransport("quadcopter")

	fmt.Println(scooter)
	fmt.Println(quad)
}