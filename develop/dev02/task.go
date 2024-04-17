package main

import (
	"container/list"
	"fmt"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" 	 => "abcd"
	- "45" 		 => "" (некорректная строка)
	- "" 		 => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const incorrectString = "(некорректная строка)"

type data struct {
	builder *strings.Builder
	stack	*list.List
	flag	bool
}

func newData() *data {
	return &data{
		builder: &strings.Builder{},
		stack: list.New(),
		flag: false,
	}
}

func (d *data) last() (rune, error) {
	list := d.stack.Back()
	if list == nil {
		return 0, fmt.Errorf("stack is empty")
	}
	defer d.stack.Remove(list)
	symb := list.Value.(rune)
	return symb, nil
}

func (d *data) repeat(n int) error {
	symb, err := d.last()
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	for i := 0; i < n; i++ {
		d.builder.WriteRune(symb)
	}
	return nil
}

func (d *data) swap(j rune) {
	if symb, ok := d.last(); ok == nil {
		d.builder.WriteRune(symb)
		d.stack.PushBack(j)
	}
}


// Unpack - функция, осуществляющая примитивную распаковку строки
func Unpack(s string) (string, error) {
	if len(s) < 1 {
		return "", nil
	}

	data := newData()
	ecran := false

	for _, j := range s {
		if j == '\\' && !ecran {
			ecran = true
			continue
		}

		if unicode.IsDigit(j) {
			if ecran {
				data.swap(j)
				ecran = false
				continue
			}
			err := data.repeat(int(j) - 48)
			if err != nil {
				return "", fmt.Errorf("%s", incorrectString)
			}
			data.flag = false
			continue
		}

		if !data.flag {
			data.stack.PushBack(j)
			data.flag = true
			continue
		} 

		data.swap(j)

		ecran = false
	}

	l, err := data.last()
	if err == nil {
		data.builder.WriteRune(l)
	}
	
	return data.builder.String(), nil
}

func main() {
	exampls := [...]string {
		"a4bc2d5e",
		"abcd",
		"45",	
		"",
		`qwe\4\5`,
		`qwe\45`,
		`qwe\\5 `,
	}

	for i := 0; i < 7; i++ {
		fmt.Println(Unpack(exampls[i]))
	}

}
