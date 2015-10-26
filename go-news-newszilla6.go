package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "newszilla6.xs4all.nl:119")
	connreader := bufio.NewReader(conn)
	message := ""

	// read in input from stdin
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Text to send: ")
	//text, _ := reader.ReadString('\n')
	// send to socket
	fmt.Fprintf(conn, "HELP\n")
	// listen for reply
	for {
		message, _ = connreader.ReadString('\n')
		fmt.Println("Message from server: " + message)
		if strings.Index(message, ".") == 0 {
			fmt.Println("End from server: " + message)
			break
		}
	}
}
