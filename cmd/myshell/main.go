package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ");
		var s, err = bufio.NewReader(os.Stdin).ReadString('\n');
		if err != nil {
			fmt.Println(err);
			return;
		}
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
                pathEnv := os.Getenv("PATH");
                paths := strings.Split(pathEnv, ":");
                found := false;
                for _, path := range paths {
                    filePath := filepath.Join(path, parts[1]);
                    if _, err := os.Stat(filePath); err == nil {
                        fmt.Print(parts[1], " is ", filePath, "\n");
                        found = true;
                        break;
                    }
                }
                if !found {
                    fmt.Print(parts[1], ": not found\n");
                }
            }
		} else {
            cmd := exec.Command(parts[0], parts[1:]...)
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            err := cmd.Run()
            if err != nil {
                fmt.Println(err)
            }
        }

	}
}
