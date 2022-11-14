package main

func BubbleSort(arr []int) []int {

	//外循环的i,表示一共只需要n-1次操作
	for i := 0; i < len(arr)-1; i++ {
		//内循环的j表示,需要进行比较的元素的索引
		//比如说一开始,i==0,那么j+1就要遍历到最后一个元素,所以j的值最大只能是len(arr)-i-1
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				//j比j+1大,那么就交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		//这个循环结束,len(arr)-i的位置就确定了,拿到了最大值,继续进行循环

	}
	return arr

}
