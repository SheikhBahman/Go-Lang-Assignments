/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/


package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func main() {
	FractionSortingSystem()
}

func FractionSortingSystem() {
	var inputSlice []int
	for i := 0; i <= 15; i++ {
		var inputString string
		fmt.Print("Enter number: ")
		fmt.Scan(&inputString)
		if appendInt, err := strconv.Atoi(inputString); err == nil {
			inputSlice = append(inputSlice, appendInt)
		} else {
			fmt.Println("Please input a number or 'X' to exit")
		}
	}
	var wg sync.WaitGroup
	wg.Add(4)
	go SortingFunction(inputSlice[:4], &wg)
	go SortingFunction(inputSlice[4:8], &wg)
	go SortingFunction(inputSlice[8:12], &wg)
	go SortingFunction(inputSlice[12:16], &wg)
	wg.Wait()
	sort.Ints(inputSlice)
	fmt.Println(inputSlice)
	FractionSortingSystem()
}

func SortingFunction(sliceToSort []int, wg *sync.WaitGroup) {
	fmt.Println(sliceToSort)
	sort.Ints(sliceToSort)
	wg.Done()
}
