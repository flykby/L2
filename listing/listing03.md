Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Так как интерфейс это структура iface, которая хранит в себе itable(в котором хранятся
методы и поля), то в err у нас будет лежать <nil>


Ответ:
<nil>
false
```
