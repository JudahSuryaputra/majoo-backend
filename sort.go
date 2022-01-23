package main

import "fmt"

func main() {
	arr := []int{
		4, -7, -5, 3, 2,
	}

	left := 0
	right := len(arr) - 1
	sortedArray := quicksort(arr, left, right)
	fmt.Println(sortedArray)
}

func quicksort(arr []int, left, right int) []int {
	if left < right {
		arr, pivot := partition(arr, left, right)
		arr = quicksort(arr, left, pivot-1)
		arr = quicksort(arr, pivot+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) ([]int, int) {
	pivot := arr[right]
	i := left

	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

/*
PSEUDOCODE

Define:
	array
	left = 0
	right = length(array) - 1

Algorithm:

	def quicksort(array, left, right):
		if left < right:
			array, pivot = partition(array, left, right)
			array = quicksort(array, left, pivot -1)
			array = quicksort(array, pivot+1, right)
		return array

	def partition(array, left, right):
		pivot = array[right] (most right element)
		i = left

		for j = left to right - 1:
			if arrary[j] < pivot:
				swap array[i] with array[j]
				i++

		swap array[i] with array[right]
*/
