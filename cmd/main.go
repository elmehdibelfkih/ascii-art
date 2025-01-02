package main

import (
	"ascii-art/internal"
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Error: number of args not valid")
		fmt.Println("Usage: go run cmd/. [STRING]")
		return
	}
	str := os.Args[1]
	internal.ParsFile("./fonts/standard.txt")
	internal.PrintStringAscii(str)
}
