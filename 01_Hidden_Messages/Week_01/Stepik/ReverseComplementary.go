package main

import (
	"fmt"
	"os"
	"bufio"
)

func ReverseComp(dna string) string {
	var Comp string = ""
	for _, r := range dna {
		if string(r) == "A" {
			Comp += "T"
		} else if string(r) == "T" {
			Comp += "A"
		} else if string(r) == "C" {
			Comp += "G"
		} else if string(r) == "G" {
			Comp += "C"
		}
	}

	var revComp string = ""
	for i := len(Comp); i > 0; i-- {
		revComp += string(Comp[i - 1])
	}

	return revComp
}

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
	str_line, _ := r.ReadString('\n')

	output := ReverseComp(str_line)

	fmt.Println(output)
}