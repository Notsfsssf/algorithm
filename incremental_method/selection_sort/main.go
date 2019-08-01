package main

import "fmt"
//选择排序
func main() {
	array := []int{5, 8, 2, 10}
	sortSlice(array[:])
	fmt.Println(array)
}

//切片
func sortSlice(array []int) {
	for i := 0; i < len(array); i++ {
		j := i + 1
		for minIndex := i; j < len(array); j++ {
			if array[i] >array[j] {
				minIndex = j
			}
			if j == len(array)-1 {
				temp := (array)[i]
				array[i] = array[minIndex]
				array[minIndex] = temp
			}
		}
	}
}
//指针
func sort(array *[]int) {
	for i := 0; i < len(*array); i++ {
		j := i + 1
		for minIndex := i; j < len(*array); j++ {
			if (*array)[i] > (*array)[j] {
				minIndex = j
			}
			if j == len(*array)-1 {
				temp := (*array)[i]
				(*array)[i] = (*array)[minIndex]
				(*array)[minIndex] = temp
			}
		}
	}

}
