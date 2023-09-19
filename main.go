package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Family interface {
	marry()
}

func (p Person) marry(){
	p.isMarried = true
}

type Person struct{
	name string
	age int
	isMarried bool
}

func main(){
	rand.Seed(time.Now().UnixNano())
	person := make([]Person, 5)
	for i:= 0; i < 5; i++ {
		person[i] = Person{
			name: "John",
			age: rand.Intn(20+30),
			isMarried: (i % 2 == 0),
		}
	}
	for _,p := range person {
		if (p.age >= 22){
			p.marry()
			fmt.Printf("%s is married \n", p.name)
		}else{
			fmt.Printf("%s is under age and cannot married \n", p.name)
		}
	}
}
