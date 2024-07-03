package main

import (
	"bufio"
	"fmt"
	"os"
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

		switch strings.Split(strings.TrimSpace(input), " ")[0]{
		case "echo":
			fmt.Fprintf(os.Stdout, strings.Join(strings.Split(input, " ")[1:], " "))
		default:
			fmt.Fprintf(os.Stdout, strings.TrimSpace(input)+": command not found\n")
		}
	
		
	}
}
