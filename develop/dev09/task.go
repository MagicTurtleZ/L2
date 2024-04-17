package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// закомментировать перед тестированием
func flagsParser(){
	flag.Parse()
}

func customWget(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("the connection to the site cannot be established: %v", err)
	}
	defer resp.Body.Close()
	
	file, err := os.Create("website.txt") 
	if err != nil {
		return fmt.Errorf("file creation error: %v", err)
	}

	if _, err = io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("file recording error: %v", err)
	}
	return nil
}

func main() {
	flagsParser()
	fmt.Println(customWget(flag.Arg(0)))
}
