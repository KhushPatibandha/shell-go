package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		var s, err = bufio.NewReader(os.Stdin).ReadString('\n');
		s = s[:len(s)-1]
		
		if s == "exit 0" {
			return
		}
		
		if err != nil {
			fmt.Println(err)
			return
		}
		
		fmt.Print(s, ": command not found\n")
	}
}
