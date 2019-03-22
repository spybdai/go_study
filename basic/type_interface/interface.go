
package main

import (
	. "fmt"
)

func main() {
	exampleOfInterface()
}

type Pet interface {	
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string
	category string
}

func (dog Dog) Name() string {
	return dog.name
}


func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Category() string {
	return dog.category
}

func exampleOfInterface() {
	dog := Dog {
		"piggy",
		"dog",
	}

	// all the following 10 ways would work	
	// value and pointers can be converted by go

	Println(dog.name)
	Println(dog.category)

	Println((&dog).name)
	Println((&dog).category)

	Println(dog.Name())
	Println(dog.Category())

	Println((&dog).Name())
	Println((&dog).Category())

	(&dog).SetName("piggy")
	Println(dog.name)

	dog.SetName("piggy 2")
	Println(dog.name)

	var pet Pet

	// &dog is an implementation of Pet
	// static type of pet is Pet
	// dynamic type of pet is *Dog
	// dynamic value of pet is &dog
	pet = &dog

	//Println(pet.name)
	//Println(pet.category)
	Println(pet.Name())
	Println(pet.Category())
	pet.SetName("piggy 3")
	Println(pet.Name())
}

