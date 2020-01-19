package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	// Get File Name as Input
	var fname string

	fmt.Println("Enter Path to File:")
	fmt.Scan(&fname)

	// Open File
	f, err := os.Open(fname)

	if err != nil {
		panic(err)
	}

	// Create Reader
	r := bufio.NewReader(f)

	// Readline
	Pattern, _ := r.ReadString('\n')
	Genome, _ := r.ReadString('\n')

	Pattern = strings.TrimSuffix(Pattern, "\n")
	Genome = strings.TrimSuffix(Genome, "\n")

	var output []int
	var substr string
	for i := 0; i < len(Genome) - len(Pattern) + 1; i++ {
		substr = Genome[i:i + len(Pattern)]
		
		if substr == Pattern {
			output = append(output, i)
		}
	}

	fmt.Println(output)
}