package main

import "fmt"

func binarySearch(arr []int,x int) int{
	left := 0
	right :=len(arr) -1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x{
			left = mid + 1
		}else{
			right = mid -1
		}
	}
	return -1
}

func main() {
	arr := []int{1,3,5,7,9,11}
	x := 7
	index := binarySearch(arr,x)
	if index != -1{
		fmt.Println("elemento trovato all'indice ",index)
	}else{
		fmt.Println("elemento non trovato")
	}
}
