package main

import (
    "fmt"
)

type Animal interface{
        Eat()
        Move()
        Speak()
}

type animalStruct struct {
	foodEaten, locomotionMethod, spokenSound string
}

func (animal animalStruct) Eat() {
	fmt.Println("it eats", animal.foodEaten)
}

func (animal animalStruct) Move() {
	fmt.Println("its locomotion is", animal.locomotionMethod)
}

func (animal animalStruct) Speak() {
	fmt.Println("it spounds", animal.spokenSound)
}


func main() {

        var animalreq, method string

        data := map[string] Animal {
		"cow"   : animalStruct{"grass", "walk", "moo"},
		"bird"  : animalStruct{"worms", "fly", "peep"},
		"snake" : animalStruct{"mice", "slither", "hsss"},
	}


        for{
		fmt.Print("please enter an animal and a method:> ")	

		fmt.Scan(&animalreq, &method)


                var animalInterface Animal


                if val, ok := data[animalreq]; ok {
                        animalInterface = val                
                } else{
                        fmt.Print("the animal not found!\n")
                        continue
                }

                 

		if method=="eat"{
			animalInterface.Eat()
		} else if method=="move"{
			animalInterface.Move()
		} else if method=="speak"{
			animalInterface.Speak()
		} else {
                    fmt.Print("method not found!\n")	    
                }
	}  

}