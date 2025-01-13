package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var Standard = make(map[rune][]string, 95)

const CharacterHeight = 8

func ParsFile(fontFileName string) {
	data, err := os.ReadFile(fontFileName)
	if err != nil {
		log.Fatalln(err)
	}
	str := strings.Trim(string(data), "\n")
	tmp := strings.Split(str, "\n\n")
	for i, j := 32, 0; i <= 126; i++ {
		Standard[rune(i)] = strings.Split(tmp[j], "\n")
		j++
	}
}

func PrintStringAscii(str string) {
	if strings.Trim(str, `\n`) == "" {
		fmt.Print(strings.Repeat("\n", strings.Count(str, `\n`)))
		return
	}
	tmp := strings.Split(str, `\n`)
	for _, line := range tmp {
		if line != "" {
			for i := 0; i < CharacterHeight; i++ {
				for _, c := range line {
					if c < 127 && c > 31 && i < len(Standard[c]) {
						fmt.Print(Standard[c][i])
					}
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}
