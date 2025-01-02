package internal

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const Content = ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_'abcdefghijklmnopqrstuvwxyz{|}~`

var Standard = make(map[rune][]string, len(Content))

const CharacterHeight = 8

func ParsFile(fontFileName string) {
	data, err := os.ReadFile(fontFileName)
	if err != nil {
		log.Fatalln(err)
	}
	str := strings.Trim(string(data), "\n")
	tmp := strings.Split(str, "\n\n")
	for i, elm := range Content {
		if i < len(tmp) {
			Standard[elm] = strings.Split(tmp[i], "\n")
		}
	}
}

func PrintStringAscii(str string) {
	tmp := strings.Split(str, `\n`)
	for _, line := range tmp {
		if line != "" {
			for i := 0; i < CharacterHeight; i++ {
				for _, c := range line {
					if strings.Contains(Content, string(c)) && i < len(Standard[c]){
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
