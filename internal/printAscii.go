package internal

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strings"
)

var Standard = make(map[rune][]string, 95)

const CharacterHeight = 8
const originalHash = "ac85e83127e49ec42487f272d9b9db8b"

func ParsFile(fontFileName string) {
	data, err := os.ReadFile(fontFileName)
	hashObj := md5.New()
	hashObj.Write(data)
	hash := fmt.Sprintf("%x", hashObj.Sum(nil))
	if string(hash) != originalHash {
		fmt.Fprintf(os.Stderr, "Error: Hash does not match!\n")
		fmt.Fprintf(os.Stderr, "Computed Hash: %s\n", hash)
		fmt.Fprintf(os.Stderr, "Expected Hash: %s\n", originalHash)
		os.Exit(0)
	}
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
