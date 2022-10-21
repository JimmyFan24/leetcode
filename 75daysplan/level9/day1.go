package level9

func EqualFrequency(word string) bool {

	m := make(map[byte]int, 26)
	for i := 0; i < len(word); i++ {
		m[word[i]]++
	}
	freq := make(map[int]int)
	maxFreq := 0
	for _, v := range m {
		freq[v]++
		if v > maxFreq {
			maxFreq = v
		}
	}
	//1. 如果全部都是1
	if maxFreq == 1 {
		return true
	}
	//2.如果频次不是2种
	if len(freq) > 2 {
		return false
	}

	//3.如果只有一种频率而且只有一个字符
	if len(freq) == 1 {
		if freq[maxFreq] == 1 {
			return true
		}
		return false
	}
	//4.如果频次是两种

	for k, v := range freq {
		//5.如果频次小的是1，而且个数是1，可以返回true
		if k == 1 && v == 1 {
			return true
		}

		if k != maxFreq {
			//找到小频次
			if maxFreq-k != 1 {
				//6.如果小频次的不是1,而且差距不是1
				return false
			} else {
				//3.如果差距是1，但是最大频次的个数超过两个
				if freq[maxFreq] > 1 {
					return false
				}
			}

		}
	}

	return true
}

func CanPlaceFlowers(flowerbed []int, n int) bool {

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if i == 0 {
			if i+1 < len(flowerbed) && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
				continue
			}
		}
		if i == len(flowerbed)-1 {
			if i-1 >= 0 && flowerbed[i-1] == 0 {
				n--
				flowerbed[i] = 1
				continue
			}
		}
		if i+1 < len(flowerbed) && i-1 >= 0 {
			if flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n == 0
}
