package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() { // func main() -- точка входа
	myBuf := &bytes.Buffer {}

	io.WriteString(myBuf,"Hello World!")

mystr, _ := io.ReadAll(myBuf)
fmt.Println(string(mystr))
}
