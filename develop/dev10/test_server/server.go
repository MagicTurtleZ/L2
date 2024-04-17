package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Создаем слушатель на порту 8081
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server started on port 8081")

	// Основной цикл сервера
	for {
		// Принимаем входящее подключение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %s\n", err)
			continue
		}
		fmt.Println("Client connected")

		// Запускаем горутину для обработки подключения
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Горутина для отправки сообщений каждые 5 секунд
	go func() {
		for {
			time.Sleep(5 * time.Second)
			message := "Hello from server"
			if _, err := conn.Write([]byte(message)); err != nil {
				fmt.Printf("Failed to send message: %s\n", err)
				return
			}
			fmt.Println("Message sent to client")
		}
	}()

	// Чтение и вывод полученных сообщений в stdout
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Failed to read from connection: %s\n", err)
			return
		}
		fmt.Printf("Received message from client: %s\n", string(buffer[:n]))
	}
}