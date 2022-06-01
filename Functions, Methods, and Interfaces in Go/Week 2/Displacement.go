package main

import (
    "fmt"
)



func GenDisplaceFn(acc float64, vel float64, dis float64) func (float64) float64{

        return func (t float64) float64{
                return 0.5 * acc * t * t + vel * t + dis
        }

}


func main() {
   

  
    var acceleration float64;
    var ivelocity float64;
    var idisplacement float64
    var time float64;

    fmt.Println("Please enter the acceleration:")
    if  _, err := fmt.Scan(&acceleration); err != nil {
            fmt.Println("  Scan failed, due to ", err)
            return
    }

    fmt.Println("Please enter the initial velocity:")
    if  _, err := fmt.Scan(&ivelocity); err != nil {
            fmt.Println("  Scan failed, due to ", err)
            return
    }

    fmt.Println("Please enter the initial displacement:")
    if  _, err := fmt.Scan(&idisplacement); err != nil {
            fmt.Println("  Scan failed, due to ", err)
            return
    }



    fn := GenDisplaceFn(acceleration, ivelocity, idisplacement)

    fmt.Println("Please enter the time:")
    if  _, err := fmt.Scan(&time); err != nil {
            fmt.Println("  Scan failed, due to ", err)
            return
    }


   fmt.Println("Displacement is: ", fn(time))

}