package level8

/*
Your friend is typing his name into a keyboard.
Sometimes, when typing a character c,
the key might get long pressed,
and the character will be typed 1 or more times.
You examine the typed characters of the keyboard.
Return True if it is possible that it was your friends name,
with some characters (possibly none) being long pressed.

Example 1:

Input: name = "alex", typed = "aaleex"
"l e e l e e"
"ll e ee l e e"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
*/
func IsLongPressedName(name string, typed string) bool {

	nameIndex := 0
	typedIndex := 0
	if len(typed) < len(name) {
		return false
	}
	for ; nameIndex < len(name); nameIndex++ {
		//这里需要注意,nameIndex肯定是不会索引越界的,但是typedIndex会
		if typedIndex >= len(typed) {
			return false
		}
		if name[nameIndex] != typed[typedIndex] {
			return false
		}
		if nameIndex+1 < len(name) && name[nameIndex+1] == name[nameIndex] {
			//如果name存在重复的,那么只需要指针+1就行
			typedIndex++
		} else {
			//如果name当前遍历的字符和下一个不一样,那么typed要把相同的全部遍历
			for typedIndex < len(typed) && typed[typedIndex] == name[nameIndex] {
				typedIndex++
			}
		}
	}
	if typedIndex < len(typed) {
		return false
	}
	return true
}

/*
Given an array arr of integers, check if there exist two indices i and j such that :
i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
*/
func CheckIfExist(arr []int) bool {

	m := make(map[int]int, len(arr))
	for i, v := range arr {
		if _, ok := m[v]; ok {
			if v == 0 {
				return true
			}
		}
		m[v] = i

	}
	for _, v := range arr {
		if _, ok := m[v*2]; ok {
			if m[v*2] != m[v] {
				//存在为0的情况,需要校验下两个数的index是否相同
				return true
			}

		}
	}
	return false
}

/*
You are given a string word that consists of digits and lowercase English letters.

You will replace every non-digit character with a space. For example, "a123bc34d8ef34" will become " 123  34 8  34". Notice that you are left with some integers that are separated by at least one space: "123", "34", "8", and "34".

Return the number of different integers after performing the replacement operations on word.

Two integers are considered different if their decimal representations without any leading zeros are different.
*/
func NumDifferentIntegers(word string) int {

	//1.0开头的字母忽略0
	integerMap := make(map[string]int)
	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			continue
		}
		tmp := ""
		q := i
		//忽略0
		for q < len(word) && word[q] == '0' {
			q++
		}
		for q < len(word) && word[q] >= '1' && word[q] <= '9' {
			tmp += string(word[q])
			q++
		}
		integerMap[tmp]++
		i = q - 1

	}
	return len(integerMap)
}

/*
Given an array points where points[i] = [xi, yi] represents a point on the X-Y plane,
return true if these points are a boomerang.

A boomerang is a set of three points that are all distinct and not in a straight line.
*/
func IsBoomerang(points [][]int) bool {

	//1.先排除同一条线
	//.横线
	if points[0][1] == points[1][1] && points[1][1] == points[2][1] {
		return false
	}
	//.竖线
	if points[0][0] == points[1][0] && points[1][0] == points[2][0] {
		return false
	}
	//0和1相等
	if points[0][0] == points[1][0] && points[0][1] == points[1][1] {
		return false
	}
	//0和2相等
	if points[0][0] == points[2][0] && points[0][1] == points[2][1] {
		return false
	}
	//1和2相等
	if points[1][0] == points[2][0] && points[1][1] == points[2][1] {
		return false
	}
	//.斜线,y2-y1/x2-x1 == y1-y0/x1-x0

	if points[2][0]-points[1][0] == 0 || points[1][0]-points[0][0] == 0 {
		return true
	}
	k1 := float64(points[2][1]-points[1][1]) / float64(points[2][0]-points[1][0])
	k2 := float64(points[1][1]-points[0][1]) / float64(points[1][0]-points[0][0])
	if k1 == k2 {
		return false
	}

	//2.如果不在一条线上,排除有两个点或者三个点相同的情况
	//三个点相同 的情况已经在横竖线那里排除了

	return true
}
