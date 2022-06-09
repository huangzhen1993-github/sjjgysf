package array

//存在重复元素2
//给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
//找出两个重复元素并且当它们的索引值差值的绝对值小于等于k时，返回true

//暴力求解
func ContainsNearbyDuplicate(nums []int,k int) bool {
	for i := 0;i < len(nums);i++{
		for j := i + 1;j < len(nums);j++{
			if nums[i] == nums[j] && (j - i) <= k{
				return true
			}
		}
	}
	return false
}

//使用map缓存数据，减少时间复杂度
func containsNearbyDuplicate2 (nums []int,k int) bool {
	//_mp := map[int]int{}
	_mp := make(map[int]int)
	for i := 0;i < len(nums);i++{
		if _v,_ok := _mp[nums[i]];_ok && (i - _v) <= k{
			return true
		}else{
			_mp[nums[i]] = i
		}
	}
	return false
}
