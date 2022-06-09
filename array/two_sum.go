package array


//两数之和实现
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//给定 nums = [2, 7, 11, 15], target = 9因为 nums[0] + nums[1] = 2 + 7 = 9所以返回 [0, 1]

//暴力求解
func TowSum (nums []int,target int) []int{
	for i := 0;i < len(nums);i++{
		for j := i + 1;j < len(nums);j++ {
			if nums[i] + nums[j] == target{
				return []int{i,j}
			}
		}
	}
	return nil
}

//借用map减少时间复杂度
func TowSum1 (nums []int,target int) []int{
	//_m := map[int]int{}
	_m := make(map[int]int)
	for i := 0;i < len(nums);i++{
		if _v,_ok := _m[target - nums[i]];_ok{
			return []int{_v,i}
		}else{
			_m[nums[i]] = i
		}
	}
	return nil
}