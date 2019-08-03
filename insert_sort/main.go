package main

import "fmt"

//插入排序
func main() {
	array := []int{492, 38, 65, 97, 487, 11, 55, 38, 54848, 88, 447}
	quickSort(array[:])
	fmt.Println(array)
}

func quickSort(array []int) {
	for i := 1; i < len(array); i++ {
		a := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > a; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = a
	}
}
