package main

import (
    "fmt"
    "strings"
)

func main() {

     
    fmt.Println("Please enter a string:")
  
    var userInput string

    fmt.Scan(&userInput)

    userInput = strings.ToLower(userInput)

    if (userInput[0:1] == "i" && 
    userInput[len(userInput) - 1:] == "n" &&
    strings.Contains(userInput, "a")){
        fmt.Print("Found!")
    } else{
        fmt.Print("Not Found!")
    }

}