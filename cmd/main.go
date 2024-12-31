package main

import (
	"fmt"
	"os"
	"ascii-art/internal"
	"ascii-art/config"
)

func main() {
	str := os.Args[1]
	// _ = str
	if len(os.Args[1:]) != 1 {
		fmt.Println("Error: no arguments")
		return
	}
	config.Standard = internal.ParsFile("./fonts/standard.txt")
	config.Shadow = internal.ParsFile("./fonts/shadow.txt")
	// config.Thinkertoy = internal.ParsFile("./fonts/thinkertoy.txt")
	internal.PrintStringAscii(str)

}