package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ");
		var s, err = bufio.NewReader(os.Stdin).ReadString('\n');
		s = s[:len(s)-1];
		
		var parts = strings.Split(s, " ");

		if parts[0] == "exit" && parts[1] == "0" && len(parts) == 2 {
			return;
		} else if parts[0] == "echo" {
			for i := 1; i < len(parts); i++ {
				if i == len(parts) - 1 {
					fmt.Print(parts[i]);
				} else {
					fmt.Print(parts[i] + " ");
				}
			}
			fmt.Print("\n")
		} else {
			fmt.Print(s, ": command not found\n");
		}

		if err != nil {
			fmt.Println(err);
			return;
		}
	}
}
