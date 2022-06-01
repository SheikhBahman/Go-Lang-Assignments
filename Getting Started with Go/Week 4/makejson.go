package main

import (
    "fmt"
    "encoding/json"
    "bufio"
     "os"
)


func main() {
    
   
   
    m := make(map[string]string)

    fmt.Println("Please enter a name:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() 
    m["name"] = scanner.Text()


    fmt.Println("Please enter the address:")   
    scanner.Scan() 
    m["address"] = scanner.Text()

    u, err := json.Marshal(m)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(u))


}