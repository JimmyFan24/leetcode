package level5

import "sort"

func CanVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool, len(rooms))
	for i := 1; i < len(rooms); i++ {
		visited[i] = false
	}
	visited[0] = true

	dfsSearchRoom(&visited, rooms, 0, len(rooms)-1, 0)
	for _, v := range visited {
		if v == false {
			return false

		}
	}
	return true

}
func dfsSearchRoom(visited *map[int]bool, rooms [][]int, count int, end int, start int) {

	if end == count {

		return
	}

	i := start
	if i >= len(rooms) || i < 0 {
		return
	}
	for j := 0; j < len(rooms[i]); j++ {
		if (*visited)[rooms[i][j]] == true {
			continue
		}
		(*visited)[rooms[i][j]] = true
		dfsSearchRoom(visited, rooms, count+1, end, rooms[i][j])
	}

	return
}

func PartitionArray(nums []int, k int) int {

	sort.Ints(nums)
	count := 0
	for len(nums) >= 1 {
		i := 1
		for i < len(nums) && isSmallerThanK(nums[i], nums[0], k) {
			i++
		}
		count++
		nums = nums[i:]
	}
	return count

}
func isSmallerThanK(max, min, k int) bool {
	return max-min <= k
}
