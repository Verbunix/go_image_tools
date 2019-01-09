package main

import (
	"fmt"
	_ "fmt"
	"os"
	_ "os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
}
