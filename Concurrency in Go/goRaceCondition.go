package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {   

    const num = 10
    var wg sync.WaitGroup
    wg.Add(num)
    counter := 0
    for i := 0; i < num; i++ {
        go func() {
            temp := counter
            runtime.Gosched()
            temp++
            counter = temp
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println("count:", counter)
}

/* Explaniation: 10 go functions that increment the same shared integer variable are launched in 10 separate goroutines;
                 each function read the varibales, increment it and write it back but as they are in a race condition 
                 the initial value of the variable in eash goroutines is not guaranteed therfore the program will print 
                 differnet values in each run.
+/ 

/*

coder@e05b5b95d459:~/project$ go run -race goRaceCondition.go 
==================
WARNING: DATA RACE
Read at 0x00c000016070 by goroutine 8:
  main.main.func1()
      /home/coder/project/goRaceCondition.go:19 +0x3c

Previous write at 0x00c000016070 by goroutine 6:
  main.main.func1()
      /home/coder/project/goRaceCondition.go:22 +0x5c

Goroutine 8 (running) created at:
  main.main()
      /home/coder/project/goRaceCondition.go:18 +0xea

Goroutine 6 (finished) created at:
  main.main()
      /home/coder/project/goRaceCondition.go:18 +0xea
==================
==================
WARNING: DATA RACE
Read at 0x00c000016070 by goroutine 7:
  main.main.func1()
      /home/coder/project/goRaceCondition.go:19 +0x3c

Previous write at 0x00c000016070 by goroutine 6:
  main.main.func1()
      /home/coder/project/goRaceCondition.go:22 +0x5c

Goroutine 7 (running) created at:
  main.main()
      /home/coder/project/goRaceCondition.go:18 +0xea

Goroutine 6 (finished) created at:
  main.main()
      /home/coder/project/goRaceCondition.go:18 +0xea
==================
count: 10
Found 2 data race(s)
exit status 66

*/