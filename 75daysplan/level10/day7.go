package level10

import "math"

func MaxNumberOfBalloons(text string) int {

	m := make(map[byte]int)
	for i := 0; i < len(text); i++ {
		m[text[i]]++
	}
	count := math.MaxInt32

	//balloon
	bal := "balloon"
	for i := 0; i < len(bal); i++ {

		if _, ok := m[bal[i]]; !ok {
			count = 0
			break
		} else {
			if bal[i] == 'n' || bal[i] == 'b' || bal[i] == 'a' {
				if m[bal[i]] <= count {
					count = m[bal[i]]
				}
			} else {
				if m[bal[i]]/2 <= count {
					count = m[bal[i]] / 2
				}
			}

		}
	}

	return count
}

func ReverseOnlyLetters(s string) string {

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if isEnglishLetter(s[i]) {
			stack = append(stack, s[i])
		}
	}
	res := ""
	index := len(stack) - 1
	for i := 0; i < len(s); i++ {
		if !isEnglishLetter(s[i]) {
			res += string(s[i])
		} else {

			if index >= 0 {
				res += string(stack[index])
			}
			index--
		}
	}

	return res
}
func isEnglishLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || (c >= 'A' && c <= 'Z')

}
