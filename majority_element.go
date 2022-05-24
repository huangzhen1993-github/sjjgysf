package main

import "sort"

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

//示例 1:
//输入：[3,2,3]
//输出：3

//示例 1:
//输入：[2,2,1,1,1,2,2]
//输出：2


//先排序
func MajorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

//使用map
func MajorityElement1(nums []int) int {
	_mp := map[int]int{}
	for _,_v := range nums{
		if _,_ok := _mp[_v];_ok{
			_mp[_v]++;
			if _mp[_v] > (len(nums)/2){
				return _v
			}
		}else{
			_mp[_v]++
		}
	}
	return 0
}

//摩尔投票法
//每次从序列里选择两个不相同的数字删除掉（或称为“抵消”），最后剩下一个数字或几个相同的数字，就是出现次数大于总数一半的那个。
//首先，可以证明最终不会一个数字都不剩。 原因： 假设两两抵消之后，最终一个数字都不剩。那么就是说一共有偶数个数字，假设有n个，那么n = 2k，k是整数。所以是进行了k次两两抵消。又因为一定存在众数 (数量超过⌊n/2⌋ = k的数字 )，所以该众数出现次数至少为k+1。由抽屉原理，这就会导致前面两两抵消的某一对数字是一样的。这是矛盾的。所以这就证明了最终不会一个数字都不剩，至少剩下一个。  假设最终剩下的那一种数字是a，假设前面进行了k次两两抵消。要证明a是欲求的众数，即证明其他数字不可能是众数。由于抽屉原理，在前面抵消的数字中，同一种数字最多出现k次，即是除了a之外的数字最多出现k次。而且最终至少剩下一个数字，所以数字的总数量大于等于2k+1。那么除了a之外的数字出现的频率<= k/(2k+1) < k/2k = 1/2，所以证明了除了a之外的数字均不会是众数。那么就是说最终剩下的那种数字a是所求众数
func MajorityElement2(nums []int) int {
	candidate, count := nums[0], 1
	for _, v := range nums {
		if candidate == v {
			count++
		} else {
			count--
			if count == 0 {
				candidate = v
				count = 1
			}
		}
	}
	return candidate
}
