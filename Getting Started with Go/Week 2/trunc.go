package main

import "fmt"

func main() {

     
    fmt.Println("Please enter a float number:")
  
    var userInput float64

    fmt.Scanln(&userInput)    
    fmt.Print("Your truncated number is: ")

    fmt.Print(int(userInput))
	fmt.Print("\n")
}