package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
*/

// сложная подсистема состоит из множества разнообразных классов
type Drum struct {
}

func (d *Drum) Play() {
	fmt.Println("Playing the drum")
}

type Trumpet struct {
}

func (t *Trumpet) Play() {
	fmt.Println("Playing the trumpet")
}

type Violin struct{
}

func (t *Violin) Play() {
	fmt.Println("Playing the violin")
}


// фасад предоставляет быстрый доступ к определённой функциональности подсистемы
type OrchestraFacade struct {
	drum 	*Drum
	trumpet *Trumpet
	violin 	*Violin
}

func NewOrchestraFacade() *OrchestraFacade {
	return &OrchestraFacade{
		drum: 		&Drum{},
		trumpet: 	&Trumpet{},
		violin: 	&Violin{},
	}
}

func (o *OrchestraFacade) PlayTogether() {
	fmt.Println("The orchestra is playing music!")
	o.drum.Play()
	o.trumpet.Play()
	o.violin.Play()
}

// клиентский код
func ExampleFacade() {
	orchestra := NewOrchestraFacade()
	orchestra.PlayTogether()
}

/*
	Плюсы:
1. изолирует клиентов от компонентов сложной подсистемы
	Минусы:
1. создание новой зависимости на основе суперобъекта порождаемого фасадом
*/

