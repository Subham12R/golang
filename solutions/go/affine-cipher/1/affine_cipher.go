package affinecipher

import (
	"errors"
	"strings"
	"unicode"
)

func modInverse(a, m int) int {
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return -1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Encode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		 return "", errors.New("a and m must be coprime")
	}

	var result strings.Builder
	count := 0

	for _, r := range text {
		r = unicode.ToLower(r)

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			continue
		}

		if count > 0 && count%5 == 0 {
			result.WriteRune(' ')
		}

		if unicode.IsDigit(r) {
			result.WriteRune(r)
			count++
			continue
		}

		x := int(r - 'a')

		y := (a*x + b) % 26

		result.WriteRune(rune(y) + 'a')
		count++
	}
	return result.String(), nil

}
func Decode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		return "", errors.New("a and b are not coprime")
	}

	inv := modInverse(a, 26)

	var result strings.Builder
	count := 0
	for _, r := range text {

		if r == ' ' {
			continue
		}

		if unicode.IsDigit(r) {
			result.WriteRune(r)
			count++
			continue
		}

		r = unicode.ToLower(r)

		y := int(r - 'a')
		x := inv * (y - b)

		x = ((x % 26) + 26) % 26
		result.WriteRune(rune('a' + x))
	}

	return result.String(), nil
}
