package main

import (
	"sort"
)

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

//排序后再查找
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0;i < len(nums) - 1;i += 2{
		if nums[i] != nums[i + 1]{
			return nums[i]
		}
	}
	return nums[len(nums) - 1]
}

//使用map
func singleNumber1(nums []int) int {
	m := map[int]bool{}
	for i := 0;i <= len(nums) - 1;i++ {
		if m[nums[i]]{
			delete(m,nums[i])
		}else{
			m[nums[i]] = true
		}
	}
	r := 0
	for v := range m {
		r = v
	}
	return r
}

//使用map，记录元素出现的次数
func singleNumber2(nums []int) int {
	m := make(map[int]int)

	for _, n := range nums {
		m[n] += 1
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -1
}

//位运算(一个数连续进行两次^运算其实就是0，0^任何数x都等于x)
func singleNumber4(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
