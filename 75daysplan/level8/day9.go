package level8

import "strings"

func ValidPalindrome(s string) bool {

	if len(s) <= 1 {
		return true
	}
	front, end := 0, len(s)-1
	for front < end {
		if s[front] != s[end] {
			return isPalindrome(s, front+1, end) || isPalindrome(s, front, end-1)
		}
		front++
		end--
	}
	return true
}
func isPalindrome(str string, front, end int) bool {

	for front < end {
		if str[front] != str[end] {
			return false
		}
		front++
		end--
	}
	return true
}

/*
For a string sequence, a string word is k-repeating if word concatenated k times is a substring of sequence.
The word's maximum k-repeating value is the highest value k where word is k-repeating in sequence.
If word is not a substring of sequence, word's maximum k-repeating value is 0.

Given strings sequence and word, return the maximum k-repeating value of word in sequence.
*/
func MaxRepeating(sequence string, word string) int {

	w := len(word)
	count := 0
	for i := 0; i <= len(sequence)-w; i++ {
		countTmp := 0
		j := i
		for j+w <= len(sequence) && sequence[j:j+w] == word {
			countTmp++
			j += w
		}

		count = max(count, countTmp)
	}
	return count
}

func CheckOnesSegment(s string) bool {

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			j := i + 1
			for j < len(s) && s[j] == '0' {
				j++
			}
			if j > i+1 {
				return false
			}

		}
	}
	return true
}
func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {

	res := 0
	for i := 0; i < len(experience); i++ {
		curEngery, curExperience := energy[i], experience[i]
		increaseEnergy, increaseExperience := 0, 0
		if initialEnergy <= curEngery {
			increaseEnergy = curEngery - initialEnergy + 1
			initialEnergy += increaseEnergy
		}
		initialEnergy -= curEngery

		if initialExperience <= curExperience {
			increaseExperience = curExperience - initialExperience + 1
			initialExperience += increaseExperience
		}
		initialExperience += curExperience

		res += increaseExperience + increaseEnergy
	}
	return res
}

func CheckStraightLine(coordinates [][]int) bool {

	//如果在一条线上，斜率相等
	x0 := coordinates[0][0]
	y0 := coordinates[0][1]
	x1 := coordinates[1][0]
	y1 := coordinates[1][1]

	//x1 != x0,因为不存在重复的点
	k := float64(y1-y0) / float64(x1-x0)
	if x1-x0 != 0 {
		k = float64(y1-y0) / float64(x1-x0)
		for i := 0; i < len(coordinates)-1; i++ {
			x2 := coordinates[i][0]
			y2 := coordinates[i][1]
			x3 := coordinates[i+1][0]
			y3 := coordinates[i+1][1]
			if x3-x2 == 0 {
				return false
			}
			if float64(y3-y2)/float64(x3-x2) != k {
				return false
			}
		}
	} else {
		//如果x1 ==x0，说明只有垂直x轴这种可能
		for i := 0; i < len(coordinates); i++ {
			if coordinates[i][0] != x0 {
				return false
			}
		}

	}
	return true
}

/*
Input: time = "?5:00"
Output: 2
Explanation: We can replace the ? with either a 0 or 1, producing "05:00" or "15:00".
Note that we cannot replace it with a 2, since the time "25:00" is invalid. In total, we have two choices.
*/
func CountTime(time string) int {

	timeStr := strings.Split(time, ":")
	hour := timeStr[0]
	minute := timeStr[1]
	hourCount := countHour(hour)
	minuteCount := countMinute(minute)
	if minuteCount == 0 {
		return hourCount
	}
	if hourCount == 0 {
		return minuteCount
	}
	return hourCount * minuteCount
}
func countHour(hour string) int {
	hourCount := 0
	if hour[0] == '?' {
		if hour[1] == '?' {
			hourCount += 24
		} else {
			if hour[1] <= '3' {
				hourCount += 3
			} else {
				hourCount += 2
			}
		}
	} else {
		if hour[1] == '?' {

			if hour[0] == '2' {
				hourCount += 4
			} else {
				hourCount += 10
			}
		}
	}
	return hourCount
}
func countMinute(minute string) int {
	minuteCount := 0

	if minute[0] == '?' {
		if minute[1] == '?' {
			minuteCount += 60
		} else {
			minuteCount += 6
		}
	} else {
		if minute[1] == '?' {
			minuteCount += 10
		}
	}
	return minuteCount
}
func ReverseVowels(s string) string {

	ori := []byte(s)
	before, after := 0, len(s)-1

	for before < after {
		for before < after && !isVowel(ori[before]) {
			before++
		}
		for before < after && !isVowel(ori[after]) {
			after--
		}
		//找到一对元音字符，交换
		if before < after {
			ori[before], ori[after] = ori[after], ori[before]
		}
		before++
		after--
	}
	res := string(ori)
	return res
}
func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func MaximumTime(time string) string {

	timeGroup := strings.Split(time, ":")
	hour := timeGroup[0]
	minute := timeGroup[1]
	return getMaxHour(hour) + getMaxMinute(minute)
}
func getMaxHour(hour string) string {

	maxHour := ""
	if hour[0] == '?' {
		if hour[1] == '?' {
			maxHour = "23"
		} else {
			if hour[1] <= '3' && hour[1] >= '0' {
				maxHour = "2" + string(hour[1])
			} else {
				maxHour = "1" + string(hour[1])
			}
		}
	} else {
		if hour[1] == '?' {
			if hour[0] == '2' {
				maxHour = "23"
			} else {
				maxHour = string(hour[0]) + "9"
			}
		} else {
			maxHour = string(hour[0]) + string(hour[1])
		}
	}
	return maxHour
}
func getMaxMinute(minute string) string {

	maxMinute := ""
	if minute[0] == '?' {
		if minute[1] == '?' {
			maxMinute = "59"
		} else {
			maxMinute = "5" + string(minute[1])
		}
	} else {
		if minute[1] == '?' {
			maxMinute = string(minute[0]) + "9"
		} else {
			maxMinute = string(minute[0]) + string(minute[1])
		}
	}
	return maxMinute
}
