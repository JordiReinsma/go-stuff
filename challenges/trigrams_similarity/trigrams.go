// Given two phrases, find the similarity of their trigrams
//
// Example:
// PhraseA: "i love your hair"
// PhraseB: "i love your long mustache"
//
// Their trigrams would be:
// TrigramA: ["i love your", "love your hair"]
// TrigramA: ["i love your", "love your long", "your long mustache"]
//
// Their similarity would be 1/4

package main

import (
	"fmt"
	"strings"
)

var exists = struct{}{}

// GetTrigrams returns a set containing the trigrams of the string
func GetTrigrams(text string) map[string]struct{} {
	// Use this if you want to compare the first or last words of a phrase
	// text = "_ " + text + " _"

	// Split the string into a list of words
	words := strings.Fields(text)

	if len(words) < 3 {
		return map[string]struct{}{}
	}

	trigrams := make(map[string]struct{}, len(words)-2)

	for i := 0; i < len(words)-2; i++ {
		trigrams[words[i]+words[i+1]+words[i+2]] = exists
	}

	return trigrams
}

// GetTrigramSimilarity returns the ratio of equal trigrams
// between two strings against all their trigrams
func GetTrigramSimilarity(a, b string) float64 {
	trigA, trigB := GetTrigrams(a), GetTrigrams(b)
	fmt.Println(trigA)
	fmt.Println(trigB)

	// Cuts out edge cases early
	if len(trigA) == 0 || len(trigB) == 0 {
		return 0.0
	}

	AplusB := len(trigA) + len(trigB)

	// A becomes A union B
	for trig := range trigB {
		trigA[trig] = exists
	}

	// AunionB is the set of all trigrams from the two strings
	AunionB := len(trigA)
	// AintersectionB is the set of the matching trigrams
	AintersectionB := AplusB - AunionB

	// Similarity is the set intersection divided by the set union
	return float64(AintersectionB) / float64(AunionB)
}

func main() {
	textA := "i love your hair"
	textB := "i love your long mustache"

	fmt.Println(GetTrigramSimilarity(textA, textB))
}
