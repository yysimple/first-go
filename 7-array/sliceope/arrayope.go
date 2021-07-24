package main

import (
	"fmt"
)

func main() {
	// 这里对应的是，类型、初始化长度、切片容量
	nums := make([]int, 3, 5)
	// nums: len = 3, cap = 5, slice3 = [0 0 0]
	fmt.Printf("nums: len = %d, cap = %d, slice3 = %v\n", len(nums), cap(nums), nums)

	// 向nums中追加一个元素
	nums = append(nums, 1)
	// nums: len = 4, cap = 5, slice3 = [0 0 0 1]
	fmt.Printf("nums: len = %d, cap = %d, slice3 = %v\n", len(nums), cap(nums), nums)

	// 再次向切片中新增元素
	nums = append(nums, 2)
	// nums: len = 5, cap = 5, slice3 = [0 0 0 1 2]
	fmt.Printf("nums: len = %d, cap = %d, slice3 = %v\n", len(nums), cap(nums), nums)

	// 再次添加元素，这个时候，容量已经小于长度了，所有切片会自动扩容为原来容量的两倍， 容量 * 2
	nums = append(nums, 3)
	// nums: len = 6, cap = 10, slice3 = [0 0 0 1 2 3]
	fmt.Printf("nums: len = %d, cap = %d, slice3 = %v\n", len(nums), cap(nums), nums)

	fmt.Println("----------")

	opeSlice()
}

// slice的操作
func opeSlice() {
	nums := []int{1, 2, 3, 4}

	// 截取0 - 2，2取不到，左闭右开的索引范围，这个跟python差不多
	nums1 := nums[0:2]
	// nums1 =  [1 2]
	fmt.Println("nums1 =", nums1)

	// 索引2 到最后
	nums2 := nums[2:]
	fmt.Println("nums2 =", nums2)

	/**
	nums = [100 2 3 4]
	nums1 = [100 2]
	这里可以看出，这种截取方式其实是浅拷贝，他们指向的是同一块引用地址
	*/
	nums1[0] = 100
	fmt.Println("nums =", nums)
	fmt.Println("nums1 =", nums1)

	// go提供的copy方法，是重新开辟一块空间
	/**
	nums = [100 2 3 4]
	nums3 = [100 123 3 4]
	*/
	nums3 := make([]int, 4)
	copy(nums3, nums)
	nums3[1] = 123
	fmt.Println("nums =", nums)
	fmt.Println("nums3 =", nums3)

}
