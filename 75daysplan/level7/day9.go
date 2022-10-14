package level7

import (
	"fmt"
	"unsafe"
)

func FlipEquiv(root1 *TreeNode, root2 *TreeNode) bool {

	if root1 == nil && root2 == nil {
		return true
	}
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	//比较叶子节点
	if root1.Val == root2.Val {

		//2个节点
		if root1.Left != nil && root1.Right != nil {
			if root2.Left != nil && root2.Right != nil {
				return (FlipEquiv(root1.Left, root2.Left) && FlipEquiv(root1.Right, root2.Right)) ||
					(FlipEquiv(root1.Left, root2.Right) && FlipEquiv(root1.Right, root2.Left))
			} else {
				return false
			}
		}

		//0个节点
		if root1.Left == nil && root1.Right == nil {
			if root2.Left == nil && root2.Right == nil {
				return true
			}
			return false
		}
		//剩下的就是一个节点的
		if root1.Left != nil && root1.Right == nil {
			if root2.Left != nil && root2.Right != nil {
				return false
			} else if root2.Left == nil && root2.Right == nil {
				return false
			} else {
				return FlipEquiv(root1.Left, root2.Left) || FlipEquiv(root1.Left, root2.Right)
			}
		} else if root1.Right != nil && root1.Left == nil {
			if root2.Left != nil && root2.Right != nil {
				return false
			} else if root2.Left == nil && root2.Right == nil {
				return false
			} else {
				return FlipEquiv(root1.Right, root2.Left) || FlipEquiv(root1.Right, root2.Right)
			}
		}
	}
	return false
}

func GetSmallestString(n int, k int) string {

	//1.要求字典序要小,所以前面每一个尽量都选a

	//n-1 <= k-c <= 26*(n-1)
	//==> k - 26*(n-1)<=c <= k-(n-1)
	//如果k - 26*(n-1) <1 只能选a
	//否则,选对应的c
	res := make([]byte, n)

	i := 0
	for n > 0 {
		lower := k - 26*(n-1)
		if lower < 1 {
			res[i] = 'a'
			k--
			i++
		} else {
			res[i] = byte('a' + lower - 1)
			k -= lower
		}
		n--

	}

	return string(res)
}

func GetSmallestString1(n int, k int) string {
	str, i, j := make([]byte, n), 0, 0
	count := 0
	for i = n - 1; i <= k-26; i, k = i-1, k-26 {
		str[i] = 'z'
		count++
	}
	if i >= 0 {
		str[i] = byte('a' + k - 1 - i)
		for ; j < i; j++ {
			str[j] = 'a'
			count++
		}
	}
	fmt.Printf("res: %T, %d", str, unsafe.Sizeof(str))
	return string(str)
}

/*
You are given two strings order and s.
All the characters of order are unique and were sorted in some custom order previously.
Permute the characters of s so that they match the order that order was sorted.
More specifically, if a character x occurs before a character y in order,
then x should occur before y in the permuted string.
Return any permutation of s that satisfies this property.
Example 2:
Input: order = "cbafg", s = "abcd"
Output: "cbad"
*/
func CustomSortString(order string, s string) string {

	m := make(map[int]byte, len(order))
	sMap := make(map[byte]int, len(s))

	for i := 0; i < len(order); i++ {
		m[i] = order[i]
	}
	for _, v := range s {
		sMap[byte(v)]++
	}

	//有顺序,想一下栈

	res := []byte{}

	for i := 0; i < len(m); i++ {
		if sMap[m[i]] > 0 {
			for sMap[m[i]] > 0 {
				res = append(res, m[i])
				sMap[m[i]]--
			}
		}
	}
	for i, _ := range sMap {
		for sMap[i] > 0 {
			res = append(res, i)
			sMap[i]--
		}

	}
	return string(res)
}
