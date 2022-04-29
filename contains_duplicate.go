package main

import "sort"
//存在重复元素
//给定一个整数数组，判断是否存在重复元素。
//如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。
//示例 1:
//输入：[1,2,3,1]
//输出：true

//暴力求解
func ContainsDuplicate (nums []int) bool{
	for i := 0;i < len(nums);i++{
		for j := i + 1;j < len(nums);j++{
			if nums[i] == nums[j]{
				return true
			}
		}
	}
	return false
}

//先排序后再判断重复
func ContainsDuplicate1 (nums []int) bool {
	sort.Sort(sort.IntSlice(nums))
	for i := 0;i < len(nums);i++{
		if i == (len(nums) - 1){
			break
		}
		if nums[i] == nums[i + 1]{
			return true
		}
	}
	return false
}

//使用map缓存数据，减少时间复杂度
func ContainsDuplicate2 (nums []int) bool {
	//_mp := map[int]bool{}
	_mp := make(map[int]bool)
	for i := 0;i < len(nums);i++{
		if _,_ok := _mp[nums[i]];_ok{
			return true
		}else{
			_mp[nums[i]] = true
		}
	}
	return false
}
