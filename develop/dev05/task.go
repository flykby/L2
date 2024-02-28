package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

*/

func main() {
	var a, b, c int
	var count, ignore, invert, fixed, lnum bool

	flag.IntVar(&a, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&b, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&c, "C", 0, "печатать ±N строк вокруг совпадения")
	flag.BoolVar(&count, "c", false, "количество строк")
	flag.BoolVar(&lnum, "n", false, "печатать номер строки")

	flag.BoolVar(&fixed, "F", false, "точное совпадение со строкой")
	flag.BoolVar(&invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&ignore, "i", false, "игнорировать регистр")

	flag.Parse()

	input := flag.Arg(0)
	substring := flag.Arg(1)

	if input == "" || substring == "" {
		panic("Пропущено имя файла или подстрока")
	}


	data := readFileToStrings(input)

	// Правило подготовки данных для последующего анализа
	prepareStrings := func(str, substr string) (string, string) {
		return str, substr
	}

	// Правило поиска подстроки в строке
	isContain := func(str, substr string) bool {
		return strings.Contains(str, substr)
	}
	
	// Правило выборки результатов по умолчанию
	handleCheck := func(check bool) bool {
		return check
	}

	if ignore {
		prepareStrings = func(str, substr string) (string, string) {
			return strings.ToLower(str), strings.ToLower(substr)
		}
	}
	if fixed {
		isContain = func(str, substr string) bool {
			return str == substr
		}
	}
	if invert {
		handleCheck = func(check bool) bool {
			return !check
		}
	}
	var conIdx []int

	for i, line := range data {
		if handleCheck(isContain(prepareStrings(line, substring))) { // если строка удовлетворяет всем условиям добавляем ее индекс в слайс
			conIdx = append(conIdx, i)
		}
	}
	switch true {
	case a > 0:
		for _, idx := range conIdx {
			for i := 0; i < a; i++ {
				if idx-a+i >= 0 {
					fmt.Println(data[idx-a+i])
				}
			}
		}
	case b > 0: // обработка ключа B
		for _, idx := range conIdx {
			for i := 0; i < b; i++ {
				if idx+i < len(data) {
					fmt.Println(data[idx+i])
				}
			}
		}
	case c > 0:
		for _, idx := range conIdx {
			for i := 0; i < c; i++ {
				if idx-c+i >= 0 {
					fmt.Println(data[idx-c+i])
				}
				if idx+i < len(data) {
					fmt.Println(data[idx+i])

				}
			}
		}
	case count:
		fmt.Println(len(conIdx))
	case lnum:
		for _, idx := range conIdx {
			fmt.Println(idx + 1)
		}
	default:
		for _, idx := range conIdx {
			fmt.Println(data[idx])
		}
	}

}

func readFileToStrings(dir string) (result []string) {
	file, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
