package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	afterFlag 	int
	beforeFlag 	int
	contextFlag int
	countFlag 	bool
	ignoreFlag 	bool
	invertFlag 	bool
	fixedFlag 	bool
	numFlag 	bool
	pattern		string
	filePath	string
)

func flagsParser() {
	flag.IntVar(&afterFlag, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&beforeFlag, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&contextFlag, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&countFlag, "c", false, "количество строк")
	flag.BoolVar(&ignoreFlag, "i", false, "игнорировать регистр")
	flag.BoolVar(&invertFlag, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&fixedFlag, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&numFlag, "n", false, "печатать номер строки")
	flag.Parse()
	formatRegular(flag.Arg(0))
	filePath = flag.Arg(1)
	if (afterFlag + beforeFlag) != 0 && contextFlag != 0 {
		contextFlag = 0
		return
	}

	if (afterFlag + beforeFlag) == 0 && contextFlag != 0 {
		afterFlag = contextFlag
		beforeFlag = contextFlag
		contextFlag = 0
	}
}

func formatRegular(s string) {
	var reg strings.Builder
	
	if ignoreFlag {
		reg.WriteString(`(?i)`)
	}

	if fixedFlag {
		reg.WriteString(regexp.QuoteMeta(s))
	} else {
		reg.WriteString(s)
	}
	
	pattern = reg.String()
}

func countMatch(reg *regexp.Regexp, scanner *bufio.Scanner) int {
	cm := 0
	for scanner.Scan() {
		m := reg.Match(scanner.Bytes())
		if m && !invertFlag {
			cm++
		}
		if !m && invertFlag {
			cm ++
		}
	}
	return cm
}

func otherMatch(reg *regexp.Regexp, scanner *bufio.Scanner) {
	lineNumber := 1
	y := false

	for scanner.Scan() {
		m := reg.Match(scanner.Bytes())
		if (m && !invertFlag) || (!m && invertFlag) {
			y = true
		}

		if y {
			output(lineNumber, scanner.Text())
		}
		lineNumber++
		y = false
	}
}

func output(lineNumber int, str string) {
	var out strings.Builder

	if numFlag {
		out.WriteString(fmt.Sprintf("%d:", lineNumber))
	}

	out.WriteString(str)

	fmt.Println(out.String())
}

func customGrep() error {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("regular expression compilation: %s", err.Error())
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("canno`t open file: %s", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if countFlag {
		fmt.Println(countMatch(reg, scanner))
		return nil
	}

	otherMatch(reg, scanner)
	return nil
}

func main() {
	flagsParser()
	customGrep()
}
