
package main

import (
	. "fmt"
)

func main() {

	exampleStruct()
}

type AnimalCategory struct {
	class string
	name  string
}

// neted with AnimalCategory
type Animal struct {
	SinceName string
	AnimalCategory
}

// 
type Cat struct {
	name string
	Animal
}

func (ac AnimalCategory) String() string {
	return Sprintf("%s-%s", ac.class, ac.name)
}

// get the unamed attribute with . operator
//func (a Animal) String() string {
//	return Sprintf("%s-%s", a.AnimalCategory, a.SinceName)
//}

// another version of method String of Animal
// can refer to method of "unnamed" attribute by "."
func (a Animal) String() string {
	return a.AnimalCategory.String()
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}

func (cat Cat) String() string {
	return Sprintf("%s-%s", cat.Animal, cat.name)
}

// pointer type of method
func (cat *Cat) Class() string {
	return  Sprintf("%s", (*cat).class)
}

// pointer type of method, version 2
// go helps with the type transaction
func (cat *Cat) Class2() string {
	return  Sprintf("%s", cat.class)
}

// method setName of *Cat, the correct version
// if passed in value, Go will help to convert to pointer
func (cat *Cat)setName(name string) {
	cat.name = name
}

// method setName of Cat, the wrong version
// the name of the origin cat object will NO change
// if passed in pointer, Go will help to convert to value
func (cat Cat)setName2(name string) {
	cat.name = name
}

func exampleStruct() {
	animalCategory := AnimalCategory {
		class: "CatClass",
		name: "Cat",
	}

	Printf("animal category is %s\n", animalCategory)

	animal := Animal {
		SinceName: "CatAnimal",
		AnimalCategory: animalCategory,
	}
	Printf("animal is %s\n", animal)
	Printf("animal category is %s\n", animal.Category())
	Printf("animal category is %s\n", animal.Category)


	cat := Cat {
		name: "catlili",
		Animal: animal,
	}	
	// cat's method "String" will be called at first priority
	Printf("cat is %s\n", cat)

	// cat has no method "Categroy", but Animal has
	Printf("cat category is %s\n", cat.Category())

	//
	Printf("cat class is %s\n", (&cat).Class())
	Printf("cat class2 is %s\n", (&cat).Class())

	// test name
	Printf("cat name is %s\n", cat.name)
	cat.setName("cat_test")
	Printf("cat name is %s\n", cat.name)
	
	(&cat).setName2("cat_test_2")
	Printf("cat name is %s\n", cat.name)
	
}

