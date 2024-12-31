package internal

import (
	"log"
	"os"
	"strings"
	"ascii-art/config"
)

func ParsFile(font string) map[rune][]string {
	ret :=  make(map[rune][]string)

	
	data, err := os.ReadFile(font)
	if err != nil {
		log.Fatalln(err)
	}
	data = []byte(strings.Trim(string(data), "\n"))
	tmp := strings.Split(string(data), "\n\n")
	// println("han>", tmp[0])
	for i, elm := range config.Content {
		// if i > len(tmp) - 1 {
		// 	break
		// }
		ret[elm] = strings.Split(tmp[i], "\n")
	}
	return ret
}
