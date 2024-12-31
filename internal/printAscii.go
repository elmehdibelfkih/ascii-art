package internal

import (
	"fmt"
	"ascii-art/config"
)

func PrintStringAscii(str string) {
	for i := 0; i < config.CharacterHeight; i++{
		for _, c := range str {
			fmt.Print(config.Standard[c][i])
		}
		fmt.Println()
	}
}