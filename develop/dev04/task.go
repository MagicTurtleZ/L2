package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func makeAnagram(word string) string {
	word = strings.ToLower(word)
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func formatMap(m map[string][]string) {
	for i, j := range m {
		if len(j) < 1 {
			delete(m, i)
		}
	}

	for i := range m {
		sort.Strings((m)[i])
	}
}

func findSets(words []string) *map[string][]string {
	anagrams := make(map[string]string)
	sets := make(map[string][]string)


	for _, word := range words {
		word = strings.ToLower(word)
		aWord := makeAnagram(word)
		if _, ok := anagrams[aWord]; !ok {
			anagrams[aWord] = word
		} else {
			sets[anagrams[aWord]] = append(sets[anagrams[aWord]], word)
		}
	}

	formatMap(sets)
	return &sets
}

func main() {
	words := []string {
		"аборт", "борат", "обрат", "табор", "торба",
		"абрек",
		"Авран", "варан", "Навар",
		"автол", "отвал",
		"автор", "втора", "отвар", "рвота", "тавро", "товар",
		"агнат", "таган",
		"агнец", "ганец",
		"амогус",
		"аграф", "графа",
		"адрес",
		"аймак", "кайма", "майка",
		"акант", "канат", "накат",
		"аксон", "накос", "носка",
		"актер", "катер", "терка",
		"актин", "антик", "нитка",
	}

	m := findSets(words)
	for i, j := range *m {
		fmt.Println(i, ": ", j)
	}
}
