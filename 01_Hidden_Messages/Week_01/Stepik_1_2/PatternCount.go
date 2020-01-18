package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// Get file name as input
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

	// DNA in the first line and pattern in the second
	byte_DNA, _, _ := r.ReadLine()
	bye_pattern, _, _ := r.ReadLine()

	string_DNA := string(byte_DNA)
	string_pattern := string(byte_pattern)

	// Count
	var count int = 0
	var substr string

	for i := 0; i < len(string_DNA) - len(string_pattern) + 1; i++ {
		substr = string_DNA[i:i + len(string_pattern)]

		if substr == string_pattern {
			count = count + 1
		}
	}

	fmt.Println("Found", count, "times")
}