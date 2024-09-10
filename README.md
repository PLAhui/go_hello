# hello
just say hello

## Install
import code
```bash
go get github.com/PLAhui/go_hello@latest
```
install cmd
````bash
go install github.com/PLAhui/go_hello@latest
````

## Example
Here's a simple example as follows:
```go
package main

import (
	"fmt"
	"github.com/PLAhui/go_hello"
)

func main() {
	result := hello.Hello("jack")
	fmt.Println(result)
}
```
