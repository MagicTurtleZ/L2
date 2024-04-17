package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type sortableLines []string

func (s sortableLines) Len() int      { return len(s) }
func (s sortableLines) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortableLines) Less(i, j int) bool {
	return compareLines(s[i], s[j], keyColumn, numericSort)
}

var (
	keyColumn   int
	numericSort bool
	reverse     bool
	unique      bool
)

// закомментировать перед тестированием  
func init() {
	flag.IntVar(&keyColumn, "k", 1, "указание колонки для сортировки")
	flag.BoolVar(&numericSort, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(filePath string, sortedLines []string) error {
	if err := os.Truncate(filePath, 0); err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range sortedLines {
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
		_, err = file.WriteString("\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func compareLines(line1, line2 string, keyColumn int, numericSort bool) bool {
	fields1 := strings.Fields(line1)
	fields2 := strings.Fields(line2)

	if keyColumn > len(fields1) || keyColumn > len(fields2) {
		return line1 < line2
	}

	var compareResult int
	if numericSort {
		num1, err1 := strconv.Atoi(fields1[keyColumn-1])
		num2, err2 := strconv.Atoi(fields2[keyColumn-1])
		if err1 != nil || err2 != nil {
			compareResult = strings.Compare(fields1[keyColumn-1], fields2[keyColumn-1])
		} else {
			compareResult = num1 - num2
		}
	} else {
		compareResult = strings.Compare(fields1[keyColumn-1], fields2[keyColumn-1])
	}

	return compareResult < 0
}

func uniqueLines(lines []string) []string {
	var uniqueLines []string
	seen := make(map[string]bool)
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}

func reverseSlice(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func customSort(filePath string) ([]string, error) {   
	lines, err := readLines(filePath)
	
	if err != nil {
		return nil, err
	}

	sort.Sort(sortableLines(lines))

	if unique {
		lines = uniqueLines(lines)
	}

	if reverse {
		reverseSlice(lines)
	}
	return lines, nil
}

func main() {
	filePath := flag.Arg(0)
	lines, err := readLines(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	sort.Sort(sortableLines(lines))

	if unique {
		lines = uniqueLines(lines)
	}

	if reverse {
		reverseSlice(lines)
	}

	if err := writeLines(filePath, lines); err != nil {
		fmt.Println(err.Error())
	}
}


