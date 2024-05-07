# golorama
A Go implementation of project colorama to get terminal colored characters

## Usage

First import the module to use:

```bash
go mod init
go import github.com/helioloureiro/golorama@latest
```

Next use it in your code like this:
```go
package main

import (
    "fmt"
    
    "github/helioloureiro/golorama"
    )


func main() {
    fmt.Println(golorama.GetCSI(golorama.GREEN) + "Hello Green World" + golorama.Reset())
    golorama.PrintColorln(golorama.RED, "Hello Red World")
}
```
