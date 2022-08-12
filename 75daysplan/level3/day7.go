package level3

func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	key := getType(ruleKey)
	for _, item := range items {
		if key != -1 {

			if ruleValue == item[key] {
				res++
			}
		}

	}
	return res
}
func getType(ruleKey string) int {

	switch ruleKey {
	case "type":
		return 0
	case "color":
		return 1
	case "name":
		return 2
	default:
		return -1

	}
}

func ArithmeticTriplets(nums []int, diff int) int {

	return 0
}
