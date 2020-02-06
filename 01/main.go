package main

import (
	"fmt"
)

func main() {}

func greet(name string) string {
	switch name {
	case "Rosie":
		return fmt.Sprintf("Hey %s, you rock!", name)

	default:
		return fmt.Sprintf("Hello %s!", name)
	}
}
