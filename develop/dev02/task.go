package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/


func Unpacking(str string) (string, error) {
	sb := strings.Builder{}
	runes := []rune(str)

	// Руна, которая будет умножать на число после нее.
	var procRune *rune

	for i := 0; i < len(runes); i++ {
		current := runes[i]

		if current == '\\' {
			if i+1 > len(runes)-1 {
				return "", errors.New("Invalid input format")
			}

			sb.WriteRune(runes[i+1])
			procRune = new(rune)
			*procRune = runes[i+1]

			i++

			// Переход на следующую итерацию.
			continue
		}

		// Если текущая руна не цифра.
		if !unicode.IsDigit(current) {
			// Запись в билдер.
			sb.WriteRune(current)

			// Установка руны как обрабатываемого.
			procRune = new(rune)
			*procRune = runes[i]
		} else {
			// Если текущая руна - цифра, то эта руна служит коэффициентом 
			// для обрабатываемой руны.
			if procRune != nil {
				// Перевод руны в число.
				count, err := strconv.Atoi(string(current))
				if err != nil {
					return "", err
				}

				// Умножаем руну N раз.
				for i := 1; i < count; i++ {
					sb.WriteRune(*procRune)
				}

				procRune = nil
			} else {
				// Если коэффициент не обабатывает руну, то строка не валидна.
				return "", errors.New("Invalid input format")
			}
		}
	}

	// Сборка строки и возврат.
	return sb.String(), nil
}