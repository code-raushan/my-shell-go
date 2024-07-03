package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("error occured while reading the shell input %v\n", err)
		}

		if strings.TrimSpace(input)=="exit 0"{
			os.Exit(0)
		}

		args := strings.Split(strings.TrimSpace(input), " ")[1:]

		switch strings.Split(strings.TrimSpace(input), " ")[0]{
		case "echo":
			result := strings.Join(strings.Split(input, " ")[1:], " ")
			fmt.Fprint(os.Stdout, result)
		case "type":
			if len(args) == 1 {
				switch args[0]{
				case "exit", "echo", "type":
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])
				default:
					paths := strings.Split(os.Getenv("PATH"), ":")

					for _, path := range paths {
						fpath := filepath.Join(path, args[0])
						if _, err := os.Stat(fpath); err != nil{
							fmt.Fprintf(os.Stdout, "%s is %s\n", args[0], fpath)
							return
						}
					}
					fmt.Fprintf(os.Stdout, "%s: not found\n", args[0])
				}
			} else {
				fmt.Fprint(os.Stdout, "too many arguments\n")
			}
		default:
			fmt.Fprintf(os.Stdout, strings.TrimSpace(input)+": command not found\n")
		}
	}
}
