package main

import (
	"fmt"
	"regexp"
	"github.com/atotto/clipboard"
	"time"
)

func main() {
	r := regexp.MustCompile(`^“([^“”]+)”[^“”]+Excerpt From:[^“”]+“[^“”]+” iBooks.`)
	for true {
		clipContent, _ := clipboard.ReadAll()
		matchedGroups := r.FindStringSubmatch(clipContent)
		if len(matchedGroups) > 0 {
			stripped := matchedGroups[1]
			clipboard.WriteAll(stripped)
			fmt.Println("\n\n" + stripped)
		}
		time.Sleep(50)
	}
}
