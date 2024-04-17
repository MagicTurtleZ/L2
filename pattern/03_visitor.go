package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
*/

// элемент описывает метод принятия посетителя, имеет единственный параметр, объявленный с типом интерфейса посетителя
type Messanger interface {
	SendHello()
	Accept(v Visitor)
}

// конкретные элементы реализуют методы принятия посетителя
type Telegram struct {
}

func (t *Telegram) SendHello() {
	fmt.Println("Hello")
}

func (t *Telegram) Accept(v Visitor) {
	v.visitForTelegram(t)
}

type WhatsApp struct {
}

func (w *WhatsApp) SendHello() {
	fmt.Println("Hello")
}

func (w *WhatsApp) Accept(v Visitor) {
	v.visitForWhatsApp(w)
}

type VK struct {
}

func (vk *VK) SendHello() {
	fmt.Println("Hello")
}

func (vk *VK) Accept(v Visitor) {
	v.visitForVK(vk)
}

// посетитель описывает общий интерфейс для всех типов посетителей
type Visitor interface {
	visitForTelegram(t *Telegram)
	visitForWhatsApp(w *WhatsApp)
	visitForVK(vk *VK)
}

// структура, реализующая конкретного визитера. Реализует методы для
// обхода конкретного элемента
type SendCustomMessage struct {
	message string
}

func (s *SendCustomMessage) visitForTelegram(t *Telegram) {
	fmt.Println(s.message)
}

func (s *SendCustomMessage) visitForWhatsApp(w *WhatsApp) {
	fmt.Println(s.message)
}

func (s *SendCustomMessage) visitForVK(vk *VK) {
	fmt.Println(s.message)
}

// клиентский код
func ExampleVisitor() {
	tg := new(Telegram)
	wp := new(WhatsApp)
	vk := new(VK)

	scm := &SendCustomMessage{message: "Hello World!"}

	tg.Accept(scm)
	wp.Accept(scm)
	vk.Accept(scm)
}

/*
	Плюсы:
1. упрощает добавление операций, работающих со сложными структурами объектов
2. объединяет родственные операции в одном классе
	Минусы:
1. паттерн не оправдан, если иерархия элементов часто меняется
*/
