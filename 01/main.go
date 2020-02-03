package main

import (
	"fmt"
)

func main() {
	fmt.Println(Greet("Everyone"))
}

func Greet(name string) string{
	return fmt.Sprintf("Hello %s!", name)
}

