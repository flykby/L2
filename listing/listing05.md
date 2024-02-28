Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
У err есть 2 параметра: тип данных и значение. Так как test() возвращает тип 
данных customError, а не пустой интерфейс, то err != nil.

Ответ: error
```
