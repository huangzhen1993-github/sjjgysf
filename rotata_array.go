package main

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//
//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右旋转 1 步: [7,1,2,3,4,5,6]
//向右旋转 2 步: [6,7,1,2,3,4,5]
//向右旋转 3 步: [5,6,7,1,2,3,4]
//示例 2:
//
//输入: [-1,-100,3,99] 和 k = 2
//输出: [3,99,-1,-100]
//解释:
//向右旋转 1 步: [99,-1,-100,3]
//向右旋转 2 步: [3,99,-1,-100]


//旋转数组
func RotataArray(nums []int,k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)//全部数组翻转
	reverse(nums, 0, k-1)//复原前k个元素
	reverse(nums, k, n-1)//复原剩下的元素
}

func reverse(nums []int, i int, j int) {
	for m, n := i, j; m < n; m, n = m+1, n-1 {
		nums[m], nums[n] = nums[n], nums[m]
	}
}

//第二种解法，新开辟一个数组
func RotataArray1(nums []int,k int) {
	n := len(nums)
	temp := make([]int, n)
	k = k % n
	for i, v := range nums {
		temp[(i+k)%n] = v
	}
	copy(nums, temp)
}