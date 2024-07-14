package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		cmd := strings.Split(strings.TrimSpace(input), " ")[0]
		args := strings.Split(strings.TrimSpace(input), " ")[1:]

		switch cmd{
		case "echo":
			result := strings.Join(strings.Split(input, " ")[1:], " ")
			fmt.Fprint(os.Stdout, result)
		case "type":
			if len(args) == 1 {
				switch args[0]{
				case "exit", "echo", "type":
					fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", args[0])
				default:
					path_var := os.Getenv("PATH")
					paths := strings.Split(path_var, ":")

					found := false

					for _, path := range paths {
						if _, err := os.Stat(path + "/" + args[0]); err == nil{
							fmt.Fprintf(os.Stdout, "%s is %s/%s\n", args[0], path, args[0])
							found = true
							break
						}
					}
					if !found {
						fmt.Fprintf(os.Stdout, "%s: not found\n", args[0])
					}
				}
			} else {
				fmt.Fprint(os.Stdout, "too many arguments\n")
			}
		default:
			// fmt.Fprintf(os.Stdout, strings.TrimSpace(input)+": command not found\n")
			command := exec.Command(args[0], args...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout

			err := command.Run()
			if err != nil {
				fmt.Fprintf(os.Stdout, cmd+": command not found\n")
			}

			// fmt.Fprintf(os.Stdout, strings.TrimSpace(input)+": command not found\n")
		}
	}
}
