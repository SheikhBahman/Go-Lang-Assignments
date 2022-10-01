/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

/*///////////////////////////////
Author: Bahman Sheikh
Date: 9/14/2022
Email: sheikh.bahman@gmail.com
Copyright © 2022 Bahman Sheikh.
//////////////////////////////*/

package main

import (
	"fmt"	
	"sort"
)

func main() {


	fmt.Println("Enter the size of array (should be at least 4)")
    var n int
	
    if m, err := Scan(&n); m != 1{		
        panic(err)
    }
	if n < 4{
		panic("Should be more than 4")
	}
    fmt.Printf("Enter %d integers\n", n)
    arrayInt := make([]int, n)
    ReadN(arrayInt, 0, n)
    fmt.Println("Entered array is:", arrayInt)

	subSize := n / 4

	channel := make(chan []int, 4)
    
	//Concurrent sort
	go sortArray(arrayInt[:subSize],                channel)
	go sortArray(arrayInt[subSize:2 * subSize],     channel)
	go sortArray(arrayInt[2 * subSize:3 * subSize], channel)
	go sortArray(arrayInt[3 * subSize:],            channel)
    
	//Recieve sorted arrays via channels
	subarray1 := <-channel; subarray2 := <-channel; subarray3 := <-channel; subarray4 := <-channel

	//Merge arrays
	arrayInt= mergeArray(mergeArray(subarray1, subarray2), mergeArray(subarray3, subarray4))
    fmt.Println("Sorted array is:", arrayInt)
}

func ReadN(all []int, i, n int) {
    if n == 0 {
        return
    }
    if m, err := Scan(&all[i]); m != 1 {
        panic(err)
    }
    ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
    return fmt.Scan(a)
}

func sortArray(array []int, channel chan []int) {
	fmt.Println("I'm sorting:", array)
	sort.Ints(array)
	channel <- array
}

func mergeArray(leftArray, rightArray []int) []int {

	mergedArray := make([]int, len(leftArray)+len(rightArray))

	i := 0
	leftIDX := 0
	rightIDX := 0

	for len(leftArray) > leftIDX && len(rightArray) > rightIDX {
		if leftArray[leftIDX] < rightArray[rightIDX] {
			mergedArray[i] = leftArray[leftIDX]
			leftIDX++
		} else {
			mergedArray[i] = rightArray[rightIDX]
			rightIDX++
		}
		i++
	}

	for j := leftIDX; j < len(leftArray); j++ {
		mergedArray[i] = leftArray[j]
		i++
	}

	for j := rightIDX; j < len(rightArray); j++ {
		mergedArray[i] = rightArray[j]
		i++
	}

	return mergedArray
}