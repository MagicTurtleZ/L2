package pattern

import (
	"fmt"
	"math/rand"
)

/*
	Реализовать паттерн «цепочка вызовов».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
*/

/*
Паттерн "«цепочка вызовов" может быть использован для реализации валидации данных,
когда необходимо проверить данные на соответствие нескольким критериям,
можно создать цепочку обязанностей, где каждый обработчик будет проверять данные на определенное условие.
Если данные проходят через все обработчики успешно, то считается что они прошли валидацию.
*/

// обработчик определяет общий для всех конкретных обработчиков интерфейс
type Validator interface {
	execute(*Request)
	setNext(Validator)
}

// структура Request, реализующая класс клиента
type Request struct {
	Login    string
	Password string
	Captcha  bool
	Role     string
}

// конкретные обработчики содержат код обработки запросов. При получении запроса каждый обработчик решает, может ли он обработать запрос,
// а также стоит ли передать его следующему объекту.
type InputFiled struct {
	next Validator
}

func (i *InputFiled) execute(r *Request) {
	if r.Login == "test" && r.Password == "123" {
		i.next.execute(r)
		return
	}

	fmt.Println("Invalid login or password!")
}

func (i *InputFiled) setNext(next Validator) {
	i.next = next
}

type Captcha struct {
	next Validator
}

func (c *Captcha) execute(r *Request) {
	b := rand.Intn(2)
	switch b {
	case 0:
		r.Captcha = false
		fmt.Println("captcha failed")
		return
	case 1:
		r.Captcha = true
		fmt.Println("captcha passed")
		c.next.execute(r)
	}
}

func (c *Captcha) setNext(next Validator) {
	c.next = next
}

type Role struct {
	next Validator
}

func (role *Role) execute(r *Request) {
	if r.Role == "user" {
		fmt.Println("successful authorisation")
	}
	fmt.Println("access denied")
}

func (role *Role) setNext(next Validator) {
	role.next = next
}

// клиентский код
func ExampleChain() {
	role := new(Role)

	captcha := new(Captcha)
	captcha.setNext(role)

	input := new(InputFiled)
	input.setNext(captcha)

	req := &Request{
		Login:    "test",
		Password: "123",
	}

	input.execute(req)
}

/*
	Плюсы:
1. уменьшает зависимость между клиентом и обработчиками.
2. реализует принцип единственной обязанности.
	Минусы:
1. есть вероятность, что запрос может остаться никем не обработанным.
*/
