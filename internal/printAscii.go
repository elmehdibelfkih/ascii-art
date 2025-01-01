package internal

import (
	"ascii-art/config"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParsFile(fontFileName string, font map[rune][]string) {
	data, err := os.ReadFile(fontFileName)
	if err != nil {
		log.Fatalln(err)
	}
	str := strings.Trim(string(data), "\n")
	println(strings.Count(str, "\n\n"))
	tmp := strings.Split(str, "\n\n")
	// println(len(tmp))
	// println(str)
	for i, elm := range config.Content {
		font[elm] = strings.Split(tmp[i], "\n")
	}
}

func PrintStringAscii(str string, font map[rune][]string) {
	tmp := strings.Split(str, `\n`)
	for _, line := range tmp {
		if line != "" {
			for i := 0; i < config.CharacterHeight; i++ {
				for _, c := range line {
					if strings.Contains(config.Content, string(c)) {
						fmt.Print(font[c][i])
					}
				}
				fmt.Println()
			}
		} else {
			fmt.Println()
		}
	}
}
