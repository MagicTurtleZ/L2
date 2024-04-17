package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type fieldsFlag []string

func (f *fieldsFlag) String() string {
	return strings.Join(*f, ", ")
}

func (f *fieldsFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

var (
	f 		 fieldsFlag
	d 		 string
	s 		 bool
	filePath string
)

func flagsParser() {
	flag.Var(&f, "f", "выбрать поля (колонки)")
	flag.StringVar(&d, "d", " ", "использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "только строки с разделителем")
	flag.Parse()
	filePath = flag.Arg(0)
}

func stringValidate(str string) {
	fields := []rune(f[0])
	sort.Slice(fields, func(i, j int) bool {return fields[i] < fields[j]})
	m := strings.Contains(str, d)
	if !m && !s {
		fmt.Println(str)
		return
	} else if !m && s {
		return
	} else {
		subs := strings.Split(str, d)
		for _, j := range fields{
			if unicode.IsDigit(j) {
				field := int(j) - 49
				if field < len(subs) {
					fmt.Printf("%s ", subs[field])
				}
			}
		}
	}
	fmt.Println()
}

func customCut() error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("canno`t open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		stringValidate(scanner.Text())
	}
	return nil
}

func stringValidateForTest(str string) ([]string, bool) {
	fields := []rune(f[0])
	sort.Slice(fields, func(i, j int) bool {return fields[i] < fields[j]})
	m := strings.Contains(str, d)
	if !m && !s {
		return []string{str}, true
	} else if !m && s {
		return nil, false
	} else {
		subs := strings.Split(str, d)
		res := make([]string, 0, len(subs))
		for _, j := range fields{
			if unicode.IsDigit(j) {
				res = append(res, subs[(int(j) - 48) - 1])
			}
		}
		return res, true
	}
}

func customCutForTest(text []string) []string {
	res := make([]string, 0, 10)
	for _, j := range text {
		if str, permit := stringValidateForTest(j); permit {
			res = append(res, str...)
		}
	}
	return res
}

func main() {
	flagsParser()
	customCut()
}
