## greeting 模块
> go mod com.meiken/module-learn/greeting
```go
package greeting

func Hello(name string) string {
	return "Hello " + name
}
```

## hello 模块
> go mod com.meiken/module-learn/hello
```go
package main

import (
	"com.meiken/module-learn/greeting"
	"fmt"
)

func main() {

	message := greeting.Hello("Gladys")
	fmt.Println(message)
}

```
> go mod tidy 

> go mod edit -replace com.meiken/module-learn/greeting=../greeting

> go mod tidy

> go run .
