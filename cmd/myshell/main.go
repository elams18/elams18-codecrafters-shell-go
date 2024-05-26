package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	// Wait for user input
	for {
		fmt.Fprint(os.Stdout, "$ ")
		s, err := bufio.NewReader(os.Stdin).ReadString('\n')
		val := strings.Split(s, "\n")
		if err == nil {
			fmt.Printf("%s: command not found\n", val[0])
		} else {
			fmt.Println(err)
		}
	}
}
