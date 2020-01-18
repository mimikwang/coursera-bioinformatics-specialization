package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type WordCount struct {
	pattern string
	count int
}

func PatternCount(Text, Pattern string) int {
	var count int = 0
	for i := 0; i < len(Text) - len(Pattern) + 1; i++ {
		if Text[i:i + len(Pattern)] == Pattern {
			count++
		}
	}

	return count
}

func Contains(s []WordCount, match WordCount) bool {
	for _, e := range s {
		if e == match {
			return true
		}
	}
	return false
}

func FrequentWords(Text string, k int) []WordCount {
	var count int = 0
	var Pattern string
	var WC WordCount
	var found bool

	FrequentPatterns := make([]WordCount, 0, 3)
	for i := 0; i < len(Text) - k; i++ {
		Pattern = Text[i:i + k]
		count = PatternCount(Text, Pattern)
		WC = WordCount{Pattern, count}

		found = Contains(FrequentPatterns, WC)
		if !found {
			FrequentPatterns = append(FrequentPatterns, WC)
		}	
	}

	count = 0
	MostPatterns := make([]WordCount, 0, 3)
	for _, wc := range FrequentPatterns {
		if wc.count > count {
			count = wc.count
		}
	}

	for _, wc := range FrequentPatterns {
		fmt.Println(wc)
		if wc.count == count {
			MostPatterns = append(MostPatterns, wc)
		}
	}

	return MostPatterns
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

	// DNA in the first line k-mer in the second
	byte_DNA, _, _ := r.ReadLine()
	byte_k, _, _ := r.ReadLine()

	string_DNA := string(byte_DNA)
	int_k, _ := strconv.Atoi(string(byte_k))

	wc := FrequentWords(string_DNA, int_k)

	fmt.Println(wc)

}