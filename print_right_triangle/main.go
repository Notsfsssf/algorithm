package main

import "fmt"

//打印直角三角形
//@@@@@
//@@@@
//@@@
//@@
//@
func main() {
	printRightTriangle(5)
}
func printRightTriangle(length int) {
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			fmt.Print("@")
			if j == length-1 {
				fmt.Print("\n")
			}
		}
	}
}
