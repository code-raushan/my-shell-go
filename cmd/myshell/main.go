package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("error occured while reading the shell input %v\n", err)
	}

	fmt.Fprintf(os.Stdout, strings.TrimSpace(input)+": command not found\n")
}
