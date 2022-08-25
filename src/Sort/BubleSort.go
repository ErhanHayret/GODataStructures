package main

import (
	"fmt"
)

func BubleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if j+1 < len(arr) && arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{0, 9, 4, 7, -1, 2, 10, 5, 1, 1, 9, 11}
	arr = BubleSort(arr)
	fmt.Println(arr)
}
