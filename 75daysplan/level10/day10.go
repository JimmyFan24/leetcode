package level10

import "strings"

func WordPattern(pattern string, s string) bool {

	subSringArrry := strings.Split(s, " ")
	if len(subSringArrry) != len(pattern) {
		return false
	}
	//Create a map with characters as keys,string as values
	patternMap := make(map[byte]string)
	stringMap := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if _, ok := patternMap[pattern[i]]; !ok {
			patternMap[pattern[i]] = subSringArrry[i]
			if _, ok := stringMap[subSringArrry[i]]; ok {
				return false
			} else {
				stringMap[subSringArrry[i]] = pattern[i]
			}

		} else {
			//fisrt,we get the value of this character key in map
			value := patternMap[pattern[i]]
			//then ,we check that whether the value is match with the current subString or not
			if value != subSringArrry[i] || stringMap[subSringArrry[i]] != pattern[i] {
				return false
			}
		}
	}
	return true
}
