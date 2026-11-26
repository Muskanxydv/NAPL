package mathutil

func Stringreverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		for _, vowel := range vowels {
			if char == vowel {
				count++
				break
			}
		}
	}
	return count
}

func Factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Power(base, exponent int) int {
	if exponent < 0 {
		return -1
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
