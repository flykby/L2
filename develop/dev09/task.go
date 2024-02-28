package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите url")

	ok := scanner.Scan()
	if !ok {
		log.Fatal("Ошибка")
	}
	url := scanner.Text()
	wget(url)
}

func wget(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Ошибка по данному url")
	}
	temp := strings.Split(url, "/")
	fileName := temp[len(temp)-1]
	saveFile, err := os.OpenFile(string(fileName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln("Ошибка при создании файла")
	}
	defer saveFile.Close()
	_, err = io.Copy(saveFile, response.Body)
	if err != nil {
		log.Fatalln("Ошибка при сохранении файла")
	}
	fmt.Println("Файл сохранен")
}
