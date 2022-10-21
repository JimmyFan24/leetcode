package weeklycontest

func AverageValue(nums []int) int {

	sum := 0
	n := 0
	for _, v := range nums {
		if v%2 == 0 && v%3 == 0 {
			sum += v
			n++
		}
	}
	if n == 0 {
		return 0
	}
	return sum / n
}

/*
Input:
creators = ["alice","bob","alice","chris"],
ids = ["one","two","three","four"],
views = [5,10,5,4]
Output: [["alice","one"],["bob","two"]]
Explanation:
The popularity of alice is 5 + 5 = 10.
The popularity of bob is 10.
The popularity of chris is 4.
alice and bob are the most popular creators.
For bob, the video with the highest view count is "two".
For alice, the videos with the highest view count are "one" and "three".
Since "one" is lexicographically smaller than "three", it is included in the answer.
*/
func MostPopularCreator(creators []string, ids []string, views []int) [][]string {

	m := make(map[string]int, len(creators))
	creatorMap := make(map[string]map[string]int, len(creators))
	res := [][]string{}
	for i, v := range creators {
		//id
		id := ids[i]
		view := views[i]
		if _, ok := creatorMap[v]; ok {
			curMap := creatorMap[v]
			curId := ""
			curView := 0
			for k, vv := range curMap {
				curView = vv
				curId = k
			}
			//分数比之前的大，直接更新
			if view > curView {
				creatorMap[v][id] = view
				delete(creatorMap[v], curId)

			} else if view == curView {
				//分数和之前的相等，看看id的大小，这里的问题在于，id并不是升序的，可能是降序的
				if curId > id {
					creatorMap[v][id] = view
					delete(creatorMap[v], curId)
				}

			}
		} else {
			mm := make(map[string]int)
			mm[id] = view
			creatorMap[v] = mm
		}

		m[v] += view
	}
	maxView := 0
	for _, v := range m {
		if v >= maxView {
			maxView = v
		}
	}
	keys := []string{}
	for _, v := range creators {
		if m[v] == maxView {
			keys = append(keys, v)
			delete(m, v)
		}
	}
	for _, v := range keys {
		id := ""
		for k, _ := range creatorMap[v] {
			id = k
		}
		res = append(res, []string{v, id})
	}

	return res
}
func MakeIntegerBeautiful(n int64, target int) int64 {

	if int(digSum(n)) <= target {
		return 0
	}

	temp, count := n, int64(10)
	for int(digSum(temp)) > target {

		mod := n % count
		temp = (n - mod) + count
		count *= 10
	}
	return temp - n
}
func digSum(n int64) int64 {
	sum := int64(0)
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}
