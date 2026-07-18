package problems

import (
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	matches := []string{}

	for a2 := range array2 {
		a2Str := array2[a2]
		for a1 := range array1 {
			a1Str := array1[a1]
			contains := strings.Contains(a2Str, a1Str)
			if contains {
				alreadyExists := false
				for _, match := range matches {
					if match == a1Str {
						alreadyExists = true
					}
				}
				if !alreadyExists {
					matches = append(matches, a1Str)
				}
			}
		}
	}

	return matches
}
