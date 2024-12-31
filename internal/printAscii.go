package internal

import (
	"ascii-art/config"
	"fmt"
	"strings"
)

func PrintStringAscii(str string) {
	tmp := strings.Split(str, `\n`)
	for _, line := range tmp {
		if line != "" {
			for i := 0; i < config.CharacterHeight; i++{
				for _, c := range line {
					fmt.Print(config.Standard[c][i])
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}