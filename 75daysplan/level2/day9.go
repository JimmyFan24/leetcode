package level2

/*
Input: matrix = [[1,5,9],[10,11,14],[12,13,15]], k = 8
Output: 13
Explanation: The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13
 */
func NewKthSmallest(matrix [][]int, k int) int {

	m,n := len(matrix),len(matrix[0])
	low,high := matrix[0][0] ,matrix[m-1][n-1]+1

	for low<high{

		mid := low + (high-low)>>1

		if counterKthSmall(m,n,mid,matrix) >=k{
			high = mid
		}else{
			low = mid+1
		}
	}
	return low
}
func counterKthSmall(m,n int,mid int,matrix[][]int)int{

	count ,j := 0,n-1
	for i:=0;i<m;i++{

		for j>=0 && mid <matrix[i][j]{
			j--
		}
		count += j+1
	}
	return count

}