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
					fmt.Print(parts[i], "\n");
				} else {
					fmt.Print(parts[i] + " ");
				}
			}
		} else if parts[0] == "type" {
			if parts[1] == "echo" || parts[1] == "exit" || parts[1] == "type" {
				fmt.Print(parts[1], " is a shell builtin\n");
			} else if parts[1] == "cat" {
				fmt.Print("cat is /bin/cat\n");
			} else {
				fmt.Print(parts[1], " not found\n");
			}
		} else {
			fmt.Print(s, ": command not found\n");
		}

		if err != nil {
			fmt.Println(err);
			return;
		}
	}
}
