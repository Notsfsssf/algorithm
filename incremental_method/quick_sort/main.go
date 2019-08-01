package main

import (
	"fmt"
)

//快速排序
func main() {
	array := []int{492, 38, 65, 97, 487, 11, 55, 38, 54848, 88, 447}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
func quickSort(array []int, left, right int) {
	if left < right {
		base := array[left]

		for left < right {
			for right > left && array[right] >= base {
				right--
			}
			array[left] = array[right]
			for right > left && array[left] <= base {
				left++
			}
			array[right] = array[left]
		}
		array[left] = base
		quickSort(array, 0, left-1)
		quickSort(array, left+1, right)
	}

}

