package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Билдер:
	Паттерн Builder определяет процесс поэтапного(в строгом порядке) построения сложного продукта.
	Он разработан для обеспечения гибкого решения различных задач создания объектов в
	объектно-ориентированном программировании. Цель шаблона проектирования builder - отделить
	построение сложного объекта от его представления.
	Применяется, когда необходимо создавать сложные объекты с большим количеством опциональных
	параметров.

Плюсы:
	1. Разделение сложного объекта
	2. Гибкость
	3. Код становится более читаемым

Минусы:
	1. Сложность добавления новых компонентов
	2. Необходимость создания отдельных строителей

Применяется:
	Применяется, когда необходимо создавать сложные объекты с большим количеством опциональных параметров


В данном коде билдер создает обьект Page, который состоит из нескольких опциональных параметров
Например: 
	убрав поле navbar - мы можем сгенерировать pdf файл на основе этой страницы, который в дальнейшем можно будет распечатать

Реализуем процесс создания страницы с помощью конструктора сайтов:
*/

// Обьект, который будем собирать
type Page struct {
	navbar string
	header string
	content string
	footer string
}

type BuilderI interface {
	SetNavbar(navbar string) BuilderI
	SetHeader(header string) BuilderI
	SetContent(content string) BuilderI
	SetFooter(footer string) BuilderI

	Build() *Page
}

type PageBuilder struct {
	navbar  string
	header  string
	content string
	footer  string
}

func NewPageBuilder() *PageBuilder {
	return &PageBuilder{}
}

func (pb *PageBuilder) SetNavbar(navbar string) BuilderI {
	pb.navbar = navbar
	fmt.Println("[PageBuilder] Создали меню")
	return pb
}
func (pb *PageBuilder) SetHeader(header string) BuilderI {
	pb.header = header
	fmt.Println("[PageBuilder] Создали заголовок")
	return pb
}
func (pb *PageBuilder) SetContent(content string) BuilderI {
	pb.content = content
	fmt.Println("[PageBuilder] Создали контент")
	return pb
}
func (pb *PageBuilder) SetFooter(footer string) BuilderI {
	pb.footer = footer
	fmt.Println("[PageBuilder] Создали завершение")
	return pb
}

// реализуем логику сборки обьекта страницы и предоставление пользователю
func (pb *PageBuilder) Build() *Page {
	return &Page {
		navbar: pb.navbar,
		header: pb.header,
		content: pb.content,
		footer: pb.footer,
	}
}


func main() {
	pageBuilder := NewPageBuilder()
	page := pageBuilder.SetNavbar("Поиск, Каталог, Личный кабинет").SetHeader("Сайт").SetContent("Товары: * * * *").SetFooter("ООО \"Сайт\"").Build()

	fmt.Println(page)

}
