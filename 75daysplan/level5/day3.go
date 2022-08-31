package level5

import (
	"sort"
	"strconv"
	"strings"
)

func FindDuplicates(nums []int) []int {

	res := []int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			res = append(res, nums[i])
		}
	}
	return res
}

func SubdomainVisits(cpdomains []string) []string {
	dominMap := make(map[string]int, len(cpdomains))

	for _, v := range cpdomains {
		countDomin(v, &dominMap)
	}
	res := []string{}
	for k, v := range dominMap {
		vv := strconv.Itoa(v)
		res = append(res, vv+" "+k)
	}
	return res
}

// 9001 discuss.leetcode.com
func countDomin(domin string, dominMap *map[string]int) {

	// 9001 discuss.leetcode.com

	d := strings.Split(domin, " ")
	c := d[0]
	count, _ := strconv.Atoi(c)
	dominList := d[1]
	for len(dominList) > 0 {

		(*dominMap)[dominList] += count
		index := strings.Index(dominList, ".")
		if index == -1 {
			break
		} else {
			dominList = dominList[index+1:]
		}

	}
}

func GoodNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}

	count := 1
	findGreater(root.Val, root.Left, &count)
	findGreater(root.Val, root.Right, &count)
	return count
}
func findGreater(max int, node2 *TreeNode, count *int) {

	if node2 == nil {
		return
	}
	if max <= node2.Val {
		*count++
		max = node2.Val
	}
	if node2.Left != nil {
		findGreater(max, node2.Left, count)
	}
	if node2.Right != nil {
		findGreater(max, node2.Right, count)
	}
	return
}
