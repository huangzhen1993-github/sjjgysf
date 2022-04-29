package main

//移动零
//题目：

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//示例:
//
//输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//说明:
//
//必须在原数组上操作，不能拷贝额外的数组。
//尽量减少操作次数。

//思路：定义两个游标i,j,i永远停留在0的位置上,j一直向右移动，当j所在位置的数不为0并且i不等于j时，则与i位置的0交换，i移动到下个为0的位置
func MoveZeros(nums []int)  {
	i := 0
	for j := 0;j < len(nums);j++{
		if nums[j] != 0{
			if i != j{
				nums[i],nums[j] = nums[j],nums[i]
			}
			i++
		}
	}
}

func MoveZeros1(nums []int)  {
	i := 0
	for j := 0;j < len(nums);j++{
		if nums[j] != 0{
			if i != j{
				nums[i] = nums[j]
				nums[j] = 0
			}
			i++
		}
	}
}

func MoveZeros2(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
