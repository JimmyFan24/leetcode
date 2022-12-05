package level10

import (
	"sort"
	"strings"
	"time"
)

func CountValidWords(sentence string) int {

	//It only contains lowercase letters, hyphens, and/or punctuation (no digits)只有小写字母，横杠，

	//There is at most one hyphen '-'. If present, it must be surrounded by lowercase characters ("a-b" is valid, but "-ab" and "ab-" are not valid).
	//There is at most one punctuation mark. If present, it must be at the end of the token ("ab,", "cd!", and "." are valid, but "a!b" and "c.," are not valid).
	count := 0
	words := strings.Split(sentence, " ")
	for _, word := range words {
		if len(word) > 0 && validWord(word) {
			count++
		}
	}
	return count
}
func validWord(word string) bool {

	//只有小写字符，-，！，。 没有数字
	hyphensCount := 0
	punctuationCount := 0
	for i := 0; i < len(word); i++ {
		//如果有大写或者数字，false
		if word[i] >= '0' && word[i] <= '9' || (word[i] >= 'A' && word[i] <= 'Z') {
			return false
		}
		//如果是-

		if word[i] == '-' {
			//-要在字母之间
			if i-1 < 0 || i+1 >= len(word) || word[i-1] < 'a' || word[i+1] < 'a' || word[i-1] > 'z' || word[i+1] > 'z' {
				return false
			}
			hyphensCount++
			if hyphensCount > 1 {
				return false
			}

		}
		if word[i] == '.' || word[i] == ',' || word[i] == '!' {
			punctuationCount++
			if i != len(word)-1 || punctuationCount > 1 {
				return false
			}

		}

	}
	return true
}

const TIME_LAYOUT = "2006-01-02 15:04:05"

func DayOfYear(date string) int {

	dateStr := strings.Split(date, "-")
	firstday, _ := time.Parse(TIME_LAYOUT, dateStr[0]+"-01-01 00:00:00")
	curday, _ := time.Parse(TIME_LAYOUT, date+" 00:00:00")
	offset := curday.Unix() - firstday.Unix()
	return int(offset / (24 * 3600))
}

func GetMaximumGenerated(n int) int {

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	res := dp[1]
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = dp[i/2]
		} else {
			dp[i] = dp[i/2] + dp[i/2+1]
		}
		res = max(res, dp[i])
	}
	return res
	//nums[2 * i] = nums[i] when 2 <= 2 * i <= n
	//nums[2 * i + 1] = nums[i] + nums[i + 1] when 2 <= 2 * i + 1 <= n
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

type Node struct {
	value int
}

type Graph struct {
	nodes []*Node          // 节点集
	edges map[Node][]*Node // 邻接表表示的无向图
}

// 增加节点
func (g *Graph) AddNode(n *Node) {

	g.nodes = append(g.nodes, n)
}

// 增加边
func (g *Graph) AddEdge(u, v *Node) {

	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*u] = append(g.edges[*u], v) // 建立 u->v 的边
	g.edges[*v] = append(g.edges[*v], u) // 由于是无向图，同时存在 v->u 的边
}
func ValidPath(n int, edges [][]int, source int, destination int) bool {

	graph := &Graph{}
	for _, edge := range edges {
		sourceNode := &Node{edge[0]}
		destinationNode := &Node{edge[1]}
		graph.AddNode(sourceNode)
		graph.AddNode(destinationNode)
		graph.AddEdge(sourceNode, destinationNode)
	}

	return true
}

func IsCovered(ranges [][]int, left int, right int) bool {

	//An integer x is covered by an interval ranges[i] = [starti, endi] if starti <= x <= endi.
	for i := left; i <= right; i++ {
		j := 0
		for ; j < len(ranges); j++ {
			if i >= ranges[j][0] && i <= ranges[j][1] {
				break
			}
		}
		if j >= len(ranges) {
			return false
		}
	}

	return true
}

func LargestSumAfterKNegations(nums []int, k int) int {

	sort.Ints(nums)
	//如果有偶数次的操作
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if nums[0] == 0 {
		return sum
	}
	if nums[0] > 0 {
		//最小不是负数 1  -1
		if k%2 == 0 {
			return sum
		}
		return sum - nums[0]*2
	}

	//有负数，求下有多少个
	j := 0
	negSum := 0
	for j < len(nums) && nums[j] < 0 {
		negSum += nums[j]
		j++
	}

	//假设有三个，现在j =3，最后的负数的索引是3-1 =2
	if j >= k {
		//把前k个负数都改为正数
		kSum := 0
		for q := 0; q < k; q++ {
			kSum += nums[q]
		}
		sum += (-kSum) * 2
	} else {
		//把前j个负数变,剩下的个数放在最大的负数
		//-3,-2,-1,1,2,3,4 j =3 k = 5
		if k%j == 0 {
			//全部负数都可以变正数
			if (k/j)%2 != 0 || (j < len(nums) && nums[j] == 0) {
				sum += (-negSum) * 2
			} else {
				if j < len(nums) && -nums[j-1] > nums[j] {
					sum += (-negSum) * 2
					if (j-k)%2 != 0 {
						sum -= nums[j] * 2
					}

				} else {
					jSum := negSum - nums[j-1]
					sum += (-jSum) * 2
				}
			}
		} else {

			//如果有0，全部负数都可以改为正数
			if j < len(nums) && nums[j] == 0 {
				sum += (-negSum) * 2
			} else {
				//否则，只能留下最后一个负数
				//如果只有一个负数
				if j <= 1 {
					if k%2 != 0 {
						sum += (-negSum) * 2
					}
				} else {

					if (k-j+1)%2 != 0 {
						sum += (-negSum) * 2
					} else {
						if j < len(nums) && -nums[j-1] > nums[j] {
							sum += (-negSum) * 2
							if (j-k)%2 != 0 {
								sum -= nums[j] * 2
							}
						} else {
							jSum := negSum - nums[j-1]
							sum += (-jSum) * 2
						}

					}
				}

			}

		}

	}

	return sum
}
