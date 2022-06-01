package main

import (
    "fmt"
)

type Animal struct {
	foodEaten, locomotionMethod, spokenSound string
}

func (v Animal) Eat() {
	fmt.Println("it eats", v.foodEaten)
}

func (v Animal) Move() {
	fmt.Println("its locomotion is", v.locomotionMethod)
}

func (v Animal) Speak() {
	fmt.Println("it spounds", v.spokenSound)
}


func main() {

        var animal, method string

        data := map[string] Animal {
		"cow"   : Animal{"grass", "walk", "moo"},
		"bird"  : Animal{"worms", "fly", "peep"},
		"snake" : Animal{"mice", "slither", "hsss"},
	}


        for{
		fmt.Print("please enter an animal and a method:> ")	

		fmt.Scan(&animal, &method)

		if method=="eat"{
			data[animal].Eat()
		} else if method=="move"{
			data[animal].Move()
		} else if method=="speak"{
			data[animal].Speak()
		} else {
                    fmt.Print("method not found!\n")	    
                }
	}  

}