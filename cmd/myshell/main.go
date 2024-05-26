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
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		val := strings.Split(s, "\n")
		if strings.EqualFold(val[0], "exit 0") {
			break;
		}
		if len(val[0])>=4 && val[0][0:4] == "echo"{
			if len(val[0]) == 4 {
				continue
			}
			if val[0][4] != ' '{
				fmt.Printf("%s: command not found\n", val[0])	
				continue
			}
			to_print := val[0][5:]
			fmt.Println(to_print)
		} else {
			fmt.Printf("%s: command not found\n", val[0])
		}
	}
}
