/*
Write two goroutines which have a race condition when executed concurrently. Explain what the race condition is and how it can occur.
*/

/*///////////////////////////////
Author: Bahman Sheikh
Date: 9/07/2022
Email: sheikh.bahman@gmail.com
Copyright © 2022 Bahman Sheikh.
//////////////////////////////*/

package main

import (
	"fmt"
	"time"
)

func logicalOne(num int) {
	for i := 0; i < num; i++ {
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("-")
	}
}

func logicalZero(num int) {
	for i := 0; i < num; i++ {
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("_")
	}
}

/* This program should print 50 pulses of 1kHz clock where
 * logical zero and logical one have the same length equal to 500 us.
 *
 * However, in most case we will get less than 50 pulses with
 * the length of some logical levels >=500us. This happens
 * because of the race conditions between logicalOne() and logicalZero()
 * goroutines. The race condition means that the order of execution
 * of the instructions of two threads/goroutines is unpredictable.
 * Even logicalOne() starts first and both logicalOne() and
 * logicalZero() have 500us sleep, Go Runtime can execute logicalZero()
 * first and/or switch between logicalOne() and logicalZero()
 * goroutines unpredictably.
 */
func main() {
	fmt.Println("\n1kHz clock diagram:")
	go logicalOne(50)
	go logicalZero(50)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("\n")
}