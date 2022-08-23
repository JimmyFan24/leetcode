package level4

import (
	"sort"
)

/*
Example 1:

Input: s = "bab", t = "aba"
Output: 1
Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.
*/
func MinSteps1347(s string, t string) int {

	sM := make(map[rune]int, 26)
	tM := make(map[rune]int, 26)
	for _, v := range s {
		sM[v]++
	}
	for _, v := range t {
		tM[v]++
	}
	count := 0
	for _, v := range s {
		if tM[v] > 0 {
			tM[v]--
		} else {
			count++
		}
	}
	return count
}

func DeckRevealedIncreasing(deck []int) []int {
	//2,13,3,11,5,17,7

	//2,3,5,7,  11,13,17
	//==> 2,3,5,7,  11,17,13
	//==>2,3,5,7,  13,11,17
	//==> 2,3,5, 7,17,13,11
	//==> 2,3,11,5,17,7,13
	//==> 2,13,3,11,5,17,7
	sort.Ints(deck)
	//11,17,13 ==> 13,17
	l := len(deck) - 1
	size := 1
	for i := len(deck) - 2; i >= 1; i-- {

		tmp := make([]int, size)
		copy(tmp, deck[i:l])
		deck[i] = deck[l]
		for j, v := range tmp {
			deck[i+j+1] = v
		}
		size++
	}
	return deck

}
func TriangularSum(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	new := []int{}
	for i := 0; i < len(nums)-1; i++ {
		tmp := 0
		if nums[i]+nums[i+1] >= 10 {
			tmp = (nums[i] + nums[i+1]) % 10
		} else {
			tmp = nums[i] + nums[i+1]
		}
		new = append(new, tmp)
	}
	return TriangularSum(new)

}
func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	list1 := getElements(root1)
	list2 := getElements(root2)
	res := []int{}
	i, j := 0, 0
	for i < len(list1) && j < len(list2) {
		if list1[i] <= list2[j] {
			res = append(res, list1[i])
			i++
		} else {
			res = append(res, list2[j])
			j++
		}
	}
	for j < len(list2) {
		res = append(res, list2[j])
		j++
	}
	for i < len(list1) {
		res = append(res, list1[i])
		i++
	}
	return res
}
func getElements(root1 *TreeNode) []int {

	res := []int{}
	if root1 == nil {
		return nil
	}
	left := getElements(root1.Left)

	for _, v := range left {
		res = append(res, v)
	}
	res = append(res, root1.Val)
	right := getElements(root1.Right)

	for _, v := range right {
		res = append(res, v)
	}
	return res

}

func MaxCoins(piles []int) int {
	sort.Ints(piles)
	count := 0
	for len(piles) >= 3 {

		count += piles[len(piles)-2]
		piles = piles[1 : len(piles)-2]
	}
	return count
}

func findTheWinner(n int, k int) int {

	if n == 1 {
		return 1
	}

	queue := []int{}
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}
	for len(queue) > 1 {

		for j := 0; j < k-1; j++ {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
		queue = queue[1:]

	}
	return queue[0]
}
