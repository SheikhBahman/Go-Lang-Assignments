/*
Implement the dining philosopher’s problem with the following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

5. The host allows no more than 2 philosophers to eat concurrently.

6. Each philosopher is numbered, 1 through 5.

7.When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a 
line by itself, where <number> is the number of the philosopher.

8/ When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a
line by itself, where <number> is the number of the philosopher.

*/

/*///////////////////////////////
Author: Bahman Sheikh
Date: 9/29/2022
Email: sheikh.bahman@gmail.com
Copyright © 2022 Bahman Sheikh.
//////////////////////////////*/

package main

import (
	"fmt"
	"sync"
)

///////////Constant variables
const numberOfPhilosophers = 5

// The host allows no more than 2 philosophers to eat concurrently-> host channel buffer
const concurrentEatCapacity = 2

// Each philosopher should eat only 3 times
const maxEat = 3
///////////

///////////Tupes
type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number          int
	leftCS, rightCS *Chopstick
}

type Host struct {
	eatPosition  chan *Philosopher
	concurrentEatCapacity int
}
///////////

///////////Functions
// The host allows no more than "chanBuff" philosophers to eat concurrently-> track host channel buffer
func (host *Host) trackNumberOfEatingPhilosopher(wg *sync.WaitGroup) {
	for {
		if len(host.eatPosition) == host.concurrentEatCapacity {
			<-host.eatPosition
			<-host.eatPosition
		}
	}
}

func (philosopher *Philosopher) Eat(host *Host, wg *sync.WaitGroup) {
	for i := 0; i < maxEat; i++ {
		//if channel has an empty capacity the philosopher can eat
		host.eatPosition <- philosopher	

		philosopher.rightCS.Lock()		
		philosopher.leftCS.Lock()

		fmt.Printf("starting to eat %v\n", philosopher.number)

		philosopher.leftCS.Unlock()
		philosopher.rightCS.Unlock()

		fmt.Printf("finishing eating %v\n", philosopher.number)			
	}
	wg.Done()
}
///////////

///////////Main
func main() {
	var wg sync.WaitGroup 
	var host Host
	host.concurrentEatCapacity = concurrentEatCapacity
	host.eatPosition = make(chan *Philosopher, concurrentEatCapacity)

	chopsticks := make([]*Chopstick, numberOfPhilosophers)
	for i := 0; i < numberOfPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// The host allows no more than 2 philosophers to eat concurrently-> track host channel buffer
	go host.trackNumberOfEatingPhilosopher(&wg)

	philosophers := make([]*Philosopher, numberOfPhilosophers)
	fmt.Printf("There are %v philosophers sharing chopsticks.\n", len(philosophers))
	for i := 0; i < numberOfPhilosophers; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%numberOfPhilosophers]}
		wg.Add(1)
		go philosophers[i].Eat(&host, &wg)
	}
	
	wg.Wait()
}