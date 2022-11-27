package level10

import "sort"

func ThousandSeparator(n int) string {

	var res string

	count := 1
	for n > 0 {
		res = string(rune(n%10+'0')) + res
		n /= 10
		if count%3 == 0 && count > 3 {
			res = "." + res
		}
		count++
	}
	return res

}

func closeStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		freq1[word1[i]-'a']++
		freq2[word2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if freq1[i] == 0 && freq2[i] > 0 {
			return false
		}
	}
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}

func BuddyStrings(s string, goal string) bool {

	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		m := make(map[byte]int)
		for i := 0; i < len(s); i++ {
			m[s[i]]++
		}
		return len(s) < len(m)
	}
	diff := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
	}
	// "abcd" results in "cbad"
	//diff[0] == 0,diff[1] ==

	return len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}
