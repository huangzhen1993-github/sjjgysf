package array

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。

//输入: [9]
//输出：[1, 0]
//
//输入: [1, 2, 3]
//输出：[1, 2, 4]

//转成数字求解
func PlusOne(nums []int) []int {
	n := 0
	j := 1
	//累加
	for i := len(nums) -1;i >= 0;i--{
		n = n + nums[i] * j
		j = j * 10
	}
	n = n + 1
	nSlice := make([]int,0)
	fSlice := make([]int,0)
	for n != 0 && n % 10 != 0{
		nSlice = append(nSlice,n % 10)
		n = int(n/10)
	}
	for k := len(nSlice) - 1;k >= 0;k--{
		fSlice = append(fSlice,nSlice[k])
	}
	return fSlice
}

//倒序循环原数组，假设每一位上的数字为 i，如果 i < 9，则不会发生进位，直接在 i + 1 然后返回
//如果 i == 9，说明发生进位，i 变成 0，继续遍历
//如果跳出循环，仍未返回，说明全都是 9，则让第一位的数字为 1 ，并且数组长度加 1，其余位全为 0

func PlusOne1(nums []int) []int{
	for i := len(nums) - 1;i >= 0;i--{
		if nums[i] == 9{
			nums[i] = 0
		}else{
			nums[i] += 1
			return nums
		}
	}
	//全部是9的情况处理
	_newNums := make([]int,len(nums) + 1)
	_newNums[0] = 1
	return _newNums
}