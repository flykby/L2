package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

func main() {
	var f int
	var d string
	var s bool

	flag.IntVar(&f, "f", 0, "fields") //объялвяем флаги
	flag.StringVar(&d, "d", "\t", "delimiter")
	flag.BoolVar(&s, "s", false, "separated")

	flag.Parse() //парсим флаги введеные пользователем

	scanner := bufio.NewScanner(os.Stdin) //объявляем сканер ввода с консоли
	fmt.Println(`После завершения ввода введите "exit"`)
	var words [][]string //срез срезов для строк
	for {
		ok := scanner.Scan() //переменная показывающая что сканер сканирует
		if !ok {             //если не сканирует
			log.Fatal(errors.New("Ошибка чтения")) //выводим ошибку
		}
		line := scanner.Text() //сохраняем строку
		if line == "exit" {    //если эта строка exit
			break //выходим из цикла
		}
		if !(s && !strings.Contains(line, d)) { //если объявляен флаг и в строке нет разделителя, то пропускаем ее
			words = append(words, strings.Split(line, d)) //в обратной ситуации сохраняем
		}
	}
	if f < 0 { //отрицательное значение не может быть
		log.Fatal(errors.New("Указан некорректныый флаг")) //поэтому выводим ошибку
	}
	if f != 0 { // если указана колонка
		var columns []string      // срез для колонок
		for _, s := range words {
			columns = append(columns, s[f]) //в каждой строке выбираем слово по индексу тем самым выбираем весь столбец
		}
		fmt.Println(columns)
	} else {
		for _, s := range words {
			for _, word := range s {
				fmt.Print(word + d)
			}
			fmt.Println("")
		}

	}
}
