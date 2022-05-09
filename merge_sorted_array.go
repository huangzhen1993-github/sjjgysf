package main

import "sort"

//给你两个有序整数数组 nums1 和 nums2，请你将 nums2
//合并到 nums1 中，使 nums1 成为一个有序数组。
//说明：
//
//初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
//你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
//
//示例：
//
//输入：
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出：[1,2,2,3,5,6]

//利用排序包
func merge(nums1 []int, m int, nums2 []int, n int){
	copy(nums1[m:],nums2)
	sort.Ints(nums1)
}

//第二种解法，新开辟一个数组，然后使用双指针
func merge1(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, len(nums1))
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			temp[k] = nums1[i]
			i++
		} else {
			temp[k] = nums2[j]
			j++
		}
		k++
	}

	if i < m {
		copy(temp[k:], nums1[i:])
	}
	if j < n {
		copy(temp[k:], nums2[j:])
	}
	copy(nums1, temp)
}

//第三种解法，倒序比较，双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	if j >= 0 {
		copy(nums1[:k+1], nums2[:j+1])
	}
}