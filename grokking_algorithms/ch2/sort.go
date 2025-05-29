package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2}
	fmt.Println(selectionSort(nums))
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0

	for index, value := range arr {
		if value < smallest {
			smallest = value
			smallest_index = index
		} 
	}
	return smallest_index
}
 
func selectionSort(arr []int) []int {
	var selectedArr []int

		for i := len(arr); i != 0; i-- { 
			fmt.Println("i=",i,"len=",len(arr),"arr=",arr)
			smallest := findSmallest(arr)
			selectedArr = append(selectedArr, arr[smallest])
			arr = append(arr[:smallest], arr[smallest+1:]...)
		}
	
	return selectedArr
}