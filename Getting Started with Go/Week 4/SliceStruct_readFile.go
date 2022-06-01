package main

import (
    "fmt"
    "bufio"   
    "os"
    "strings"
)

type name struct {
    fname string
    lname string
}


func main() {

    names := make([]name, 0, 3)

    fmt.Println("Please enter the file address:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() 
    fileAddress := scanner.Text()    

    file, err := os.Open(fileAddress)
    if (err == nil){
        fmt.Println("--------------> Reading from", fileAddress)
        fscanner := bufio.NewScanner(file)        
        for fscanner.Scan() {            
            splited := strings.Split(fscanner.Text(), " ")
            
            var aName name
            aName.fname, aName.lname = splited[0], splited[1]
            names = append(names, aName)
        }

        fmt.Println("--------------> Read Successfully")

        fmt.Println("--------------> Print the slice struct")

        for _, ns := range names{
            fmt.Println(ns.fname, ns.lname)
        }

    } else{
        fmt.Println("File not found!", fileAddress)
    }
    
   
}