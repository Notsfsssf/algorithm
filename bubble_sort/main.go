package main

import "fmt"

//冒泡排序
func main() {
	array := []int{492, 38, 65, 97, 487, 11, 55, 38, 54848, 88, 447}
	bubbleSort(array[:])
	fmt.Println(array)
}

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				temp := array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}
}
