Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

```
Ответ:
Вывод программы: error
Это происходит, поскольку функция test возвращает указатель на тип customError, который в последующем
преобразуется в интерфейс error, таким образом переменная err будет содержать данные о типе customError 
и <nil> поле с данными.
