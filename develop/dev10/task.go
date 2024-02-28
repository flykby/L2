package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeoutF := flag.Int("timeout", 10, "Timeout")
	flag.Parse()

	if os.Args[1] != "go-telnet" {
		panic("Неправильный запуск")
	}

	log.Println(os.Args)

	// Парсит имя хоста и порт.
	host := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]

	// Инициализирует канал для транслирования сигналов от опереционной системы.
	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, os.Kill)

	// Преобразовывает время в нужный формат и конкатенирует имя хоста с портом.
	timeout := time.Duration(*timeoutF) * time.Second
	address := net.JoinHostPort(host, port)

	// Создается соединение.
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		panic(err)
	}

	log.Println("[+] Connected")

	// Ждёт сигнала от операционной системы. После получения, завершается работа программы.
	go func() {
		<-exitCh

		log.Println("[-] Exit")

		// Обязвтельно закрывается сокет.
		conn.Close()
		os.Exit(0)
	}()

	// Прослушивает файловый дескриптор stdin. Запись в сокет.
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Прослушивает файловый дескриптор stdout. Чтение из сокета.
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)

		os.Exit(1)
	}
}
