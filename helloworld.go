package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	file := os.Args[1]

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	for _, name := range strings.Split(string(content), "\n") {
		count[name]++
	}

	for line, amount := range count {
		if amount > 1 {
			fmt.Printf("%d\t%s\n", amount, line)
		}
	}

}
