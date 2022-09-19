package level6

func SumEvenAfterQueries(nums []int, queries [][]int) []int {

	sum := 0
	for _, v := range nums {

		if v%2 == 0 {
			sum += v
		}
	}

	res := []int{}
	for _, v := range queries {

		index := v[1]
		//如果原来是偶数
		if nums[index]%2 == 0 {

			//原来是偶数，加了之后还是偶数，那么增加这个增量就行==> 2+4 ==6 -2+4 ==2 6 +(-4)
			if (nums[index]+v[0])%2 == 0 {

				sum += v[0]
				res = append(res, sum)
				nums[index] += v[0]
			} else if (nums[index]+v[0])%2 != 0 {
				sum -= nums[index]
				res = append(res, sum)
				nums[index] += v[0]
			}
		} else {
			//如果原来是奇数
			//1.加了之后变偶数
			if (nums[index]+v[0])%2 == 0 {

				sum += v[0] + nums[index]
				res = append(res, sum)
				nums[index] += v[0]
			} else if (nums[index]+v[0])%2 != 0 {
				res = append(res, sum)
				nums[index] += v[0]
			}

		}
	}
	return res
}
