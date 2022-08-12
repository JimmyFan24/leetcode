package main


func rescursionSort(arr []int,L int,R int)  int{

	if L == R{
		return 0
	}
	mid := L + ((R - L) >> 1)

	return rescursionSort(arr,L,mid) + rescursionSort(arr,mid+1,R) + mergeSum(arr,L,mid,R)



}
func mergeSum(arr []int,L int,M int,R int) int{
	p1 := L
	p2 := M + 1
	i := 0
	sum := 0
	help := make([]int,R-L+1)

	for{
		if p1 <= M && p2 <= R{
			if arr[p1] < arr[p2]{
				help[i] = arr[p1]
				sum += arr[p1]*(R-p2+1)
				i++
				p1++
			}else {
				help[i] = arr[p2]
				i++
				p2++
			}

		}else {
			break
		}

	}

	for{
		if p1 <= M{
			help[i] = arr[p1]
			i++
			p1++
		}else {
			break
		}

	}

	for {
		if p2 <= R{
			help[i] = arr[p2]
			p2++
			i++
		}else {
			break
		}
	}

	for i,a := range help{
		arr[L+i] = a
	}
	return sum
}
func merge(arr []int,L int,M int,R int){
	p1 := L
	p2 := M + 1
	i := 0
	help := make([]int,R-L+1)

	for{
		if p1 <= M && p2 <= R{
			if arr[p1] <= arr[p2]{
				help[i] = arr[p1]
				i++
				p1++
			}else {
				help[i] = arr[p2]
				i++
				p2++
			}

		}else {
			break
		}

	}

	for{
		if p1 <= M{
			help[i] = arr[p1]
			i++
			p1++
		}else {
			break
		}

	}

	for {
		if p2 <= R{
			help[i] = arr[p2]
			p2++
			i++
		}else {
			break
		}
	}

	for i,a := range help{
		arr[L+i] = a
	}
}