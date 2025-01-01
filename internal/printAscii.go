package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const Content = ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_'abcdefjhijklmnopqrstuvwxyz{|}~`

var Standard = make(map[rune][]string, len(Content))

const CharacterHeight = 8

func ParsFile(fontFileName string, font map[rune][]string) {
	data, err := os.ReadFile(fontFileName)
	if err != nil {
		log.Fatalln(err)
	}
	str := strings.Trim(string(data), "\n")
	tmp := strings.Split(str, "\n\n")
	for i, elm := range Content {
		font[elm] = strings.Split(tmp[i], "\n")
	}
}

func PrintStringAscii(str string, font map[rune][]string) {
	tmp := strings.Split(str, `\n`)
	for _, line := range tmp {
		if line != "" {
			for i := 0; i < CharacterHeight; i++ {
				for _, c := range line {
					if strings.Contains(Content, string(c)) {
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
