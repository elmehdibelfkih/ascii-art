package main

import (
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) != 2 {
		fmt.Println("Error: number of arguments")
		return
	}
	// inp
}