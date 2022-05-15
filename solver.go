package main

import (
	"regexp"
)

// Transform is an interface for
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

	regexp, _ := regexp.Compile(`.*III.*`)
	match_indices := regexp.FindAllStringIndex(input, -1)

	for _, index_pair := range match_indices {
		start := index_pair[0]
		end := index_pair[1]
		transforms = append(transforms, input[:start]+"U"+input[end:])
	}

	return transforms
}

// Any "UU" in the input can be replaced with ""
type Transform4 string

func (t Transform4) apply(input string) []string {
	transforms := make([]string, 0)

	regexp, _ := regexp.Compile(`.*UU.*`)
	match_indices := regexp.FindAllStringIndex(input, -1)

	for _, index_pair := range match_indices {
		start := index_pair[0]
		end := index_pair[1]
		transforms = append(transforms, input[:start]+input[end:])
	}

	return transforms
}
