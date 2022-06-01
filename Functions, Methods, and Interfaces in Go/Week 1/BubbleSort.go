package main

import (
    "fmt"
)

func swap(data []int, i int) []int {
     data[i], data[i + 1] = data[i + 1], data[i]
     return data
}

func bubbleSort(data []int) []int {

    for i := 0; i < len(data); i++ {
        
        for j := 0; j < len(data) - 1; j++ {
            if data[j] > data[j + 1]{
              swap(data, j)              
            }
        }
    }
     return data
}


func main() {
   

    var numbers []int   
    var userInt int

    
    for i := 0; i < 10; i++ {
        fmt.Println("Integers entered so dar:", len(numbers))
        fmt.Println("Please enter an integer:")
        if  _, err := fmt.Scan(&userInt); err != nil {
            fmt.Println("  Scan for i failed, due to ", err)
            break
        }

        numbers = append(numbers, userInt)
    }


    fmt.Println("Number entered before sort:")
    fmt.Println(numbers)
    bubbleSort(numbers)
    fmt.Println(numbers)




}