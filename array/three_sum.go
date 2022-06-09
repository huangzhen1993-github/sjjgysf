package array

import (
	"sort"
)

//    给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//
//        注意：答案中不可以包含重复的三元组。
//
//        例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//        满足要求的三元组集合为：
//        [
//          [-1, 0, 1],
//          [-1, -1, 2]
//        ]

//暴力求解
func ThreeSum(nums []int) [][]int{
	_finalSlice := make([][]int,0)
	//防止重复使用map
	_map := map[[3]int]bool{}
	for i := 0;i < len(nums) - 2;i++{
		for j := i + 1;j < len(nums) - 1;j++ {
			for k := j + 1;k < len(nums);k++ {
				if nums[i] + nums[j] == -nums[k]{
					_slice := []int{nums[i],nums[j],nums[k]}
					sort.Ints(_slice)
					var _array [3]int
					copy(_array[:],_slice)
					if !_map[_array]{
						_finalSlice = append(_finalSlice,_slice)
						_map[_array] = true
					}
				}
			}
		}
	}
	return _finalSlice
}

//排序后，双指针解法
func ThreeSum1(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			//排序后默认nums[i]需要小于0，不然三数之和肯定不等于0
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			//防止重复
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			if nums[j]+nums[k] == -nums[i] {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					//防止重复
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					//防止重复
					k--
				}
			} else if nums[j]+nums[k] < -nums[i] {
				//较小数的游标递增
				j++
			} else {
				//较大数的游标递减
				k--
			}
		}
	}
	return res
}