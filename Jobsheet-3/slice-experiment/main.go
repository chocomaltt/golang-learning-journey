package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	fmt.Println("Isi awal: ", nums)
	fmt.Println("Len sebelum append: ", len(nums), "Cap sebelum append: ", cap(nums))

	nums = append(nums, 4)
	fmt.Println("Setelah append: ", nums)
	fmt.Println("Len setelah append: ", len(nums), "Cap setelah append: ", cap(nums))

	sub := nums[1:3]
	sub[0] = 99
	fmt.Println("sub:", sub)
	fmt.Println("nums setelah sub diubah: ", nums)
	fmt.Println("Len sub: ", len(sub), "Cap sub: ", cap(sub))

	aman := make([]int, len(sub))
	copy(aman, sub)
	aman[0] = 0
	fmt.Println("aman: ", aman, "sub tetap: ", sub)
	fmt.Println("Len aman: ", len(aman), "Cap aman: ", cap(aman))
}