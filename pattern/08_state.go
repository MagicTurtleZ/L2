package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
*/

/*
Разработка приложения простой системы управления заказами.
У заказа может быть несколько состояний: "новый", "обрабатывается", "отгружен" и "доставлен".
*/

// cостояние описывает общий интерфейс для всех конкретных состояний
type State interface {
	ProcessOrder()
}

// контекст хранит ссылку на объект состояния и делегирует ему часть работы, зависящей от состояний
type Order struct {
	newState        State
	processingState State
	shippedState    State
	deliveredState  State

	currentState State
}

func NewOrder() *Order {
	order := new(Order)

	newState := &NewState{stateOrder: order}
	processingStat := &ProcessingState{stateOrder: order}
	shippedState := &ShippedState{stateOrder: order}
	deliveredState := &DeliveredState{stateOrder: order}

	order.setState(newState)
	order.newState = newState
	order.processingState = processingStat
	order.shippedState = shippedState
	order.deliveredState = deliveredState

	return order
}

func (o *Order) setState(state State) {
	o.currentState = state
}

func (o *Order) ProcessOrder() {
	o.currentState.ProcessOrder()
}

// конкретные состояния реализуют поведения, связанные с определённым состоянием контекста.
// состояния имеют обратную ссылку на объект контекста. Через неё осуществляется смена его состояния.
type NewState struct {
	stateOrder *Order
}

func (n *NewState) ProcessOrder() {
	fmt.Println("Processing order...")
	n.stateOrder.setState(n.stateOrder.processingState)
}

type ProcessingState struct {
	stateOrder *Order
}

func (p *ProcessingState) ProcessOrder() {
	fmt.Println("Order is being processed...")
	p.stateOrder.setState(p.stateOrder.shippedState)
}

type ShippedState struct {
	stateOrder *Order
}

func (s *ShippedState) ProcessOrder() {
	fmt.Println("Order has been shipped...")
	s.stateOrder.setState(s.stateOrder.deliveredState)
}

type DeliveredState struct {
	stateOrder *Order
}

func (d *DeliveredState) ProcessOrder() {
	fmt.Println("Order has been delivered.")
	d.stateOrder.setState(d.stateOrder.newState)
}

// клиентский код
func ExampleState() {
	order := NewOrder()

	for i := 0; i < 4; i++ {
		order.ProcessOrder()
	}
}

/*
	Плюсы:
1. избавляет от множества больших условных операторов машины состояний
2. концентрирует в одном месте код, связанный с определённым состоянием
	Минусы:
1. может неоправданно усложнить код, если состояний мало и они редко меняются
*/
