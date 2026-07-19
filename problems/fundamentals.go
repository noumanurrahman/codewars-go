package problems

import (
	"fmt"
	"math"
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

func getDigitsMath(n int) []int {
	// Handle negative numbers
	if n < 0 {
		n = -n
	}

	// Edge case for zero
	if n == 0 {
		return []int{0}
	}

	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}

func SumDigPow(a, b uint64) []uint64 {
	result := []uint64{}
	for i := a; i <= b; i++ {
		fmt.Println("Checking", i)
		digits := getDigitsMath(int(i))
		var sum float64 = 0
		fmt.Println("Digits:", digits)

		for index := range digits {
			num := math.Pow(float64(digits[index]), float64(index+1))
			sum += num
		}

		if sum == float64(i) {
			result = append(result, uint64(i))
		}
	}
	return result
}
