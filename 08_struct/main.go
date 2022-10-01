package main

import ("fmt"
		"strconv")

type Person struct{
	firstName string
	lastName string
	age int
	gender string
}

func (p Person) greet() string{
	 return "Welcome "+p.lastName+" you are : "+ strconv.Itoa(p.age)+" years old"
}

func (p *Person) getMarried(name string){
	if p.gender == "m"{

	}else if p.gender == "f"{
		p.lastName = name
	}
}

func main(){
	person1 := Person{"Abiodun", "Akinlawon",24,"f"}
	person1.getMarried("Oladimeji")
	// person1.greet
	fmt.Println(person1.greet())

}