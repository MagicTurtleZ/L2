package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
*/

/*
Паттерн "Команда" может быть использован для реализации системы управления транзакциями в базе данных.
Каждая команда может представлять собой операцию базы данных, такую как "Добавить запись", "Изменить запись", "Удалить запись" и так далее.
Такая система позволяет отменять или повторять транзакции.
*/

// команда описывает общий для всех конкретных команд интерфейс
type Command interface {
	execute()
}

// конкретные команды реализуют различные запросы, следуя общему интерфейсу команд
type AddDBCommand struct {
	srg *StorageCommand
}

func (a *AddDBCommand) execute() {
	_ = a.srg.DB
	fmt.Println("process of adding a record to database")
}

type UpdateDBCommand struct {
	srg *StorageCommand
}

func (u *UpdateDBCommand) execute() {
	_ = u.srg.DB
	fmt.Println("the process of updating a database record")
}

type DeleteDBCommand struct {
	srg *StorageCommand
}

func (d *DeleteDBCommand) execute() {
	_ = d.srg.DB
	fmt.Println("database record deletion process")
}

// получатель содержит бизнес-логику программы
type StorageCommand struct {
	DB struct{}
}

func (s *StorageCommand) Insert() Command {
	return &AddDBCommand{
		srg: s,
	}
}

func (s *StorageCommand) Update() Command {
	return &UpdateDBCommand{
		srg: s,
	}
}

func (s *StorageCommand) Delete() Command {
	return &DeleteDBCommand{
		srg: s,
	}
}

// отправитель хранит ссылку на объект команды и обращается к нему, когда нужно выполнить какое-то действие. 
// Отправитель работает с командами только через их общий интерфейс
type Invoker struct {
	Commands []Command
}

func (i *Invoker) executeCommands() {
	for _, c := range i.Commands {
		c.execute()
	}
}


// клиентский код
func ExampleCommand() {
	s := StorageCommand{
		DB: struct{}{},
	}

	tasks := []Command{
		s.Insert(),
		s.Update(),
		s.Insert(),
		s.Delete(),
	}

	invoker := &Invoker{
		Commands: tasks,
	}

	invoker.executeCommands()
}

/*
	Плюсы:
1. убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
2. позволяет реализовать простую отмену и повтор операций
3. позволяет реализовать отложенный запуск операций
4. позволяет собирать сложные команды из простых
	Минусы:
1. необходимость создания множества дополнительных классов для реализации паттерна команды
*/
