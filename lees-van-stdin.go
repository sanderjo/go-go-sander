package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n\r") // only at the end of the line!
		if text == "STOP" {
			break
		}
		fmt.Println("ingetypt:", text)
	}

	fmt.Println("Klaar!")
}
