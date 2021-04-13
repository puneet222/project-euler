package main

import (
	"fmt"
	"strings"
)

var ones = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var elevens = map[int]string{
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

var tens = map[int]string {
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

const (
	THOUSAND = "thousand"
	HUNDRED = "hundred"
	AND = "and"
)

func getWords(n int) []string {
	var words []string
	// thousands
	th := n/1000
	if th > 0 {
		words = append(words, ones[th])
		words = append(words, THOUSAND)
	}
	n = n%1000 // decrease the number
	// hundreds
	h := n/100
	if h > 0 {
		words = append(words, ones[h])
		words = append(words, HUNDRED)
		if n%100 > 0 {
			// more than 100 therefore add AND
			words = append(words, AND)
		}
	}
	n = n%100 // decrease the number
	// tens
	t := n/10
	if t == 1 {
		words = append(words, elevens[n])
		n = 0 // done
	} else if t > 1 {
		words = append(words, tens[t*10])
		n = n%10
	}
	// ones
	if n > 0 {
		words = append(words, ones[n])
	}
	return words
}

func main() {
	limit := 1000
	var sum = 0
	for i := 1; i <= limit; i++ {
		words := getWords(i)
		joinWords := strings.Join(words, "")
		sum += len(joinWords)
	}
	fmt.Printf("Number letter count till %v is %v", limit, sum)
}