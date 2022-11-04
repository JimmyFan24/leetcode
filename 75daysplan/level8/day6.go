package level8

import (
	"math"
	"strconv"
)

func OddString(words []string) string {
	difference := make(map[string]int, len(words))
	//这里的技巧是把数组变成字符串当成key,不然数组是存不了的
	for _, word := range words {
		difference[strHash(word)]++
	}

	for _, v := range words {
		if difference[strHash(v)] == 1 {
			return v
		}
	}
	return ""
}
func strHash(word string) string {

	res := ""
	for i := 0; i < len(word)-1; i++ {
		c := 0
		if word[i+1] < word[i] {
			c = int(word[i] - word[i+1])
			cc := "-" + strconv.Itoa(c)
			res += cc + ","
		} else {
			c = int(word[i+1] - word[i])
			res += strconv.Itoa(c) + ","
		}

	}
	return res
}
func TwoEditWords(queries []string, dictionary []string) []string {
	res := []string{}

	for _, query := range queries {
		if canTwoEdit(query, dictionary) {
			res = append(res, query)
		}
	}
	return res
}
func canTwoEdit(query string, dictionary []string) bool {
	for _, dir := range dictionary {
		if twoEdit(query, dir) {
			return true
		}
	}
	return false
}
func twoEdit(str1 string, str2 string) bool {

	count := 2
	for i := 0; i < len(str1); i++ {

		if str1[i] != str2[i] {
			count--
		}
	}
	if count < 0 {
		return false
	}
	return true
}

func DestroyTargets(nums []int, space int) int {

	//模 space 相同的值可以被作为同一组目标摧毁，输入值是这一组目标中值最小的数字
	modMap, cntMap := make(map[int]int, len(nums)), make(map[int]int, len(nums))
	for _, v := range nums {
		mod := v % space
		cntMap[mod]++

		if cntMap[mod] == 1 {
			modMap[mod] = v
		} else if modMap[mod] > v {
			modMap[mod] = v
		}
	}

	ans := math.MaxInt32
	tmp := 0
	for k, v := range cntMap {
		//找到mod组里面最多值的
		//tmp代表当前的mod的数量
		if tmp < v || (tmp == v && modMap[k] < ans) {
			tmp = v
			ans = modMap[k]
		}
	}
	return ans
}
