package level9

import "sort"

func CheckRecord(s string) bool {

	//The student was absent ('A') for strictly fewer than 2 days total.
	attendanceRecord := make(map[byte]int)
	for i := 0; i < len(s); i++ {

		attendanceRecord[s[i]]++
	}
	for i := 0; i < len(s)-2; i++ {
		if s[i] == 'L' && s[i+1] == 'L' && s[i+2] == 'L' {
			return false
		}
	}
	if attendanceRecord['A'] >= 2 {
		return false
	}
	return true
	//The student was never late ('L') for 3 or more consecutive days.
}
func FindLengthOfLCIS(nums []int) int {

	stack := []int{}
	res := 0
	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 {
			stack = append(stack, nums[i])
		} else {
			//栈顶元素一定是i的前一个
			top := stack[len(stack)-1]
			if nums[i] > top {
				stack = append(stack, nums[i])
			} else {
				//前一个比后面的小或者相等，计算最大值，清空stack
				res = max(res, len(stack))
				//当前元素变新的栈底元素
				stack = []int{nums[i]}
			}
		}
	}
	return res
}
func SecondHighest(s string) int {

	digitsMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digitsMap[s[i]]++
		}
	}
	digits := []int{}
	for k, _ := range digitsMap {
		n := int(k - '0')
		digits = append(digits, n)
	}
	if len(digits) < 2 {
		return -1
	}
	sort.Ints(digits)
	return digits[len(digits)-2]
}
func ModifyString(s string) string {

	if len(s) == 1 && s[0] == '?' {
		return "a"
	}
	var characters = []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] != '?' {
			res += string(s[i])
		} else {
			//？在第一
			if i-1 < 0 && i+1 < len(s) {
				for _, v := range characters {
					if s[i+1] != v {
						res += string(v)
						break
					}
				}
				//?在最后
			} else if i-1 >= 0 && i+1 >= len(s) {
				for _, v := range characters {
					if s[i-1] != v {
						res += string(v)
						break
					}
				}
				//？在中间
			} else {
				pre := s[i-1]
				if pre != '?' {
					next := s[i+1]
					i := 0
					for i < len(characters) && (characters[i] == pre || characters[i] == next) {
						i++
					}
					res += string(characters[i])
				} else {
					next := s[i+1]
					i := 0
					for i < len(characters) && (characters[i] == pre || characters[i] == next) && characters[i] != res[len(res)-1] {
						i++
					}
					res += string(characters[i])
				}

			}
		}
	}

	return res
}
