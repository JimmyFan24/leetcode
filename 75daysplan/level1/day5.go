package level1

func KidsWithCandies(candies []int, extraCandies int) []bool {
	resList := make([]bool,len(candies))

	for i:=0;i<len(candies);i++{
		iCandy := candies[i] + extraCandies
		resList[i] = true
		for j:=0;j<len(candies);j++{

			if iCandy < candies[j]{
				resList[i] = false
				break
			}

		}


	}
	return resList
}
func min(a int,b int)int{
	if a <=b {
		return a

	}
	return b
}

