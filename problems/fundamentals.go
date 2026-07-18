package problems

import (
	"strings"
)

func MakeNegative(num int) int {
	// Already negative
	if num < 0 {
		return num
	}
	return 0 - num
}

func SmallestIntegerFinder(numbers []int) int {
	lowest := numbers[0]
	for _, num := range numbers {
		if num < lowest {
			lowest = num
		}
	}
	return lowest
}

func GetSize(w, h, d int) [2]int {
	var result [2]int
	result[0] = 2 * (w*h + h*d + w*d)
	result[1] = w * h * d
	return result
}

func MultipleOfIndex(ints []int) []int {
	var result []int
	for index := 1; index < len(ints); index++ {
		num := ints[index]
		if num%index == 0 {
			result = append(result, num)
		}
	}
	return result
}

func ToWeirdCase(str string) string {
	result := ""
	words := strings.Split(str, " ")
	for i, word := range words {
		for i := range word {
			char := string(word[i])
			if i%2 == 0 {
				result += strings.ToUpper(char)
			} else {
				result += strings.ToLower(char)
			}
		}
		if i != len(words)-1 {
			result += " "
		}
	}
	return result
}

func SumDigits(number int) int {
	sum := 0
	digits := []int{0}

	temp := number

	if temp == 0 {
		return 0
	}

	if temp < 0 {
		temp = -temp
	}

	for temp > 0 {
		digit := temp % 10
		digits = append(digits, digit)
		temp /= 10
	}

	for _, dig := range digits {
		sum += dig
	}

	return sum
}

func DigitalRoot(n int) int {
	result := 0
	sum := SumDigits(n)
	if sum > 9 {
		result = DigitalRoot(sum)
	} else {
		result = sum
	}
	return result
}

func MinMax(arr []int) [2]int {
	lowest := arr[0]
	highest := arr[0]

	for _, price := range arr {
		if price < lowest {
			lowest = price
		}
		if price > highest {
			highest = price
		}
	}

	return [2]int{lowest, highest}
}
