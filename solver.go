package main

import (
	"fmt"
	"strings"
)

// Transform is an interface for any sort of string transformation.
type Transform interface {
	// Returns a list of all possible strings that result from applying this transformation.
	// string must be made up of only M, I, or U
	apply(input string) []string
}

// If a string ends in I, you can add U to the end. The string here will just be the name of the transform.
type Transform1 string // Q for reviewer: what type to use here, given it doesn't matter?

func (t Transform1) apply(input string) []string {
	if len(input) > 0 && input[len(input)-1] == 'I' {
		return []string{input + "U"}
	}
	return []string{}
}

// If you have Mx, you can get Mxx
type Transform2 string

func (t Transform2) apply(input string) []string {
	if len(input) > 1 && input[0] == 'M' {
		return []string{input + input[1:]}
	}
	return []string{}
}

// Any "III" in the input can be replaced with "U"
type Transform3 string

func (t Transform3) apply(input string) []string {
	transforms := make([]string, 0)

	pattern := "III"

	// Look for the pattern at each starting character, to cover the case of overlapping matches.
	searchStart := 0
	for {
		substrMatchStart := strings.Index(input[searchStart:], pattern)
		if substrMatchStart == -1 {
			return transforms
		}
		matchStart := substrMatchStart + searchStart
		transforms = append(transforms, input[:matchStart]+"U"+input[matchStart+len(pattern):])
		searchStart = matchStart + 1
	}
}

// Any "UU" in the input can be replaced with ""
type Transform4 string

func (t Transform4) apply(input string) []string {
	transforms := make([]string, 0)

	pattern := "UU"

	// Look for the pattern at each starting character, to cover the case of overlapping matches.
	searchStart := 0
	for {
		substrMatchStart := strings.Index(input[searchStart:], pattern)
		if substrMatchStart == -1 {
			return transforms
		}
		matchStart := substrMatchStart + searchStart

		resultStr := input[:matchStart] + input[matchStart+len(pattern):]
		if !contains(transforms, resultStr) {
			transforms = append(transforms, resultStr)
		}

		searchStart = matchStart + 1
	}
}

// Toy transform - allows us to replaces 4 Is with a single U, so we can get to a correct answer
type Transform5 string

func (t Transform5) apply(input string) []string {
	transforms := make([]string, 0)

	pattern := "IIII"

	// Look for the pattern at each starting character, to cover the case of overlapping matches.
	searchStart := 0
	for {
		substrMatchStart := strings.Index(input[searchStart:], pattern)
		if substrMatchStart == -1 {
			return transforms
		}
		matchStart := substrMatchStart + searchStart
		transforms = append(transforms, input[:matchStart]+"U"+input[matchStart+len(pattern):])
		searchStart = matchStart + 1
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	var t1 Transform1 = "1" // "End in I, add U"
	var t2 Transform2 = "2" // "Mx -> Mxx"
	var t3 Transform3 = "3" // "*III* -> *U*"
	var t4 Transform4 = "4" // "*UU* -> **"
	var t5 Transform5 = "5" // "*IIII* -> *U*"

	ts := []Transform{t1, t2, t3, t4, t5}
	bag := map[string]string{"MI": ""}

	queue := make([]string, 0)
	queue = append(queue, "MI")

	stopCount := 10000
	iterCount := 0
	for {
		iterCount += 1

		if path, ok := bag["MU"]; ok {
			fmt.Printf("Soln: %v\n", path)
			break
		}
		if iterCount >= stopCount {
			fmt.Printf("No success in %v...\n", stopCount)
			break
		}
		// Pop from the queue
		str := queue[0]
		queue = queue[1:]

		// Apply all transformation to the popped string
		for tIndex, t := range ts {
			results := t.apply(str)
			for _, result := range results {
				// Only add results to the queue that haven't already been obtained in previous steps.
				if _, ok := bag[result]; !ok {
					bag[result] = bag[str] + fmt.Sprintf("%v", tIndex+1) + "->"
					queue = append(queue, result)
				}
			}
		}
	}
}
