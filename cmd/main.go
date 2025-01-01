package main

import (
	"fmt"
	"os"
	"ascii-art/internal"
	"ascii-art/config"
)

func main() {
	str := os.Args[1]
	if len(os.Args[1:]) != 1 {
		fmt.Println("Error: no arguments")
		return
	}
	internal.ParsFile("./fonts/standard.txt", config.Standard)
	internal.PrintStringAscii(str, config.Standard)

	// internal.ParsFile("./fonts/shadow.txt", config.Shadow)
	// internal.PrintStringAscii(str, config.Shadow)

// 	internal.ParsFile("./fonts/thinkertoy.txt", config.Thinkertoy)
// 	internal.PrintStringAscii(str, config.Thinkertoy)
}
