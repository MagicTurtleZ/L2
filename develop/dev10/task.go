package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	host 		string
	timeoutFlag string
	timeout		time.Duration
)

func flagsParser() {
	flag.StringVar(&timeoutFlag, "timeout", "10s", "таймаут на подключение к серверу")
	flag.Parse()
	makeHost(flag.Arg(0), flag.Arg(1))
	makeTimeout(timeoutFlag)
}

func makeHost(h, p string) {
	var builder strings.Builder
	builder.WriteString(h)
	builder.WriteRune(':')
	builder.WriteString(p)

	host = builder.String()
}

func makeTimeout(timeStr string) {
	timeout, _= time.ParseDuration(timeStr)
}

func customTelnet() error {
	conn, err := net.Dial("tcp", host)

	if err != nil {
		time.Sleep(timeout)
		return fmt.Errorf("conection server error")
	}
	defer conn.Close()
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	gs := make(chan os.Signal, 1)
	signal.Notify(gs, syscall.SIGINT, syscall.SIGTERM)
	
	wChan := make(chan []byte)
	rChan := make(chan []byte)
	wg.Add(3)
	go connectionReader(ctx, &wg, conn, rChan, gs)
	go connectionWriter(ctx, &wg, wChan)
	
	go func() {
		defer wg.Done()
		timer := time.NewTimer(timeout)
		for {		
			select {
			case <-timer.C: 
				gs <- syscall.SIGTERM
			case send := <-wChan:
				if _, err := conn.Write(send); err != nil {
					gs <- syscall.SIGTERM
					continue
				}
			case message := <-rChan:
				fmt.Println(string(message))
			case <-gs:
				conn.Close()
				cancel()
				return
			}
			
		}
	}()
	
	wg.Wait()
	close(wChan)
	close(rChan)
	return nil
}

func connectionReader(ctx context.Context, wg *sync.WaitGroup, conn net.Conn, rChan chan []byte, gs chan os.Signal) {
	defer wg.Done()
	buf := make([]byte, 1024)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				gs <- syscall.SIGTERM
				return
			}
			rChan <- buf[:n]
		}
	}
}

func connectionWriter(ctx context.Context, wg *sync.WaitGroup, wChan chan []byte) {
	defer wg.Done()
	sender := bufio.NewScanner(os.Stdin)
	go func() {
		for sender.Scan() {
			wChan <- sender.Bytes()
		}
	}()
	<-ctx.Done()
}

func main() {
	flagsParser()
	customTelnet()
}
