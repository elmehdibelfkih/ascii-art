package main

import (
	"ascii-art/config"
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
	internal.ParsFile("./fonts/standard.txt", config.Standard)
	internal.PrintStringAscii(str, config.Standard)

	// internal.ParsFile("./fonts/shadow.txt", config.Shadow)
	// internal.PrintStringAscii(str, config.Shadow)

	// internal.ParsFile("./fonts/thinkertoy.txt", config.Thinkertoy)
	// internal.PrintStringAscii(str, config.Thinkertoy)
}
