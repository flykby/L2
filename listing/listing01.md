Что выведет программа? Объяснить вывод программы.

```go
package main

import (
    "fmt"
)

func main() {
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4]
    fmt.Println(b)
}
```

Ответ:
```
Создается слайс на основе среза из массива. 
Слайс b ссылается на часть исходного массива a с индексами [1:4).

Ответ: [77, 78, 79]

```
