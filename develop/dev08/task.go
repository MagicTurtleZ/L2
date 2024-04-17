package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func cmdCd(path string) error {
	dir := path
	err := os.Chdir(dir)
	if err != nil {
		return err
	}

	return nil
}

func cmdPwd() error {
	current, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(current)
	return nil
}

func cmdEcho(args []string) {
	for _, text := range args {
		fmt.Println(text)
	}
}

func cmdKill(args []string) error {
	pid, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	blackball, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = blackball.Kill()
	if err != nil {
		return err
	}

	return nil
}

func cmdPs() error {
	cmd := exec.Command("cmd.exe", "/c tasklist")
	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

func pipe(cmds []string) error {
	var err error
	for _, cmd := range cmds {
		args := strings.Split(cmd, " ")
		switch args[0] {
		case "cd": err = cmdCd(args[1])
		case "pwd": err = cmdPwd()
		case "echo": cmdEcho(args[1:])
		case "kill": err = cmdKill(args[1:])
		case "ps": err = cmdPs()
		case `\q`: err = fmt.Errorf("exit")
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func currentPath() error {
	usr, err := user.Current()
	if err != nil {
		return err
	}

	fullPath, err := os.Getwd()
	if err != nil {
		return err
	}
	fio := strings.Split(usr.Username, `\`)
	dir := strings.Replace(fullPath, usr.HomeDir, "", -1)
	fmt.Printf("%s@%s ~%s\n$ ", fio[1], fio[0], dir)
	return nil
}

func customOC() error {
	input := bufio.NewScanner(os.Stdin)
	
	for {
		currentPath()
		if !input.Scan() {
			return input.Err()
		}
		line := input.Text()
		cmds := strings.Split(line, " | ")
		if err := pipe(cmds); err != nil {
			return err
		}
	}
}

func main() {
	fmt.Println(customOC())
}

