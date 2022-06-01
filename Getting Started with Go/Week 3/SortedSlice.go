package main

import (
    "fmt"
    "strconv"
    "sort"
)

func main() {


    var userInput string

    sortedSlice := make([]int, 0, 3)


    for{
        fmt.Println("Please enter an integer or X to exit:")
        fmt.Scan(&userInput)
        if (userInput == "X"){
            break
        }
        intVar, err := strconv.Atoi(userInput)

        if (err != nil){
            fmt.Println("Input is not valid:")
            continue
        } else{

             sortedSlice = append(sortedSlice, intVar)

             sort.Ints(sortedSlice)
             fmt.Println( sortedSlice)
            
        }
    }



}