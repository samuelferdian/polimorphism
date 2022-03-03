package main

import (
	"fmt"
)

type Animal interface {
	GetName() string
	GetSound() string
}

type Dog struct {
	Name  string
	Sound string
}

func (d Dog) GetName() string {
	return d.Name
}

func (d Dog) GetSound() string {
	return d.Sound
}

type Cat struct {
	Name  string
	Sound string
}

func (c Cat) GetName() string {
	return c.Name
}

func (c Cat) GetSound() string {
	return c.Sound
}

type Cow struct {
	Name  string
	Sound string
}

func (c Cow) GetName() string {
	return c.Name
}

func (c Cow) GetSound() string {
	return c.Sound
}

func ListAnimal(a []Animal) {
	for _, animal := range a {
		fmt.Printf("Name: %s\n", animal.GetName())
		fmt.Printf("Sound: %s\n\n", animal.GetSound())
	}
}

func main() {
	dog := Dog{
		Name:  "Echo",
		Sound: "Bark!! Bark!!",
	}

	cat := Cat{
		Name:  "Kitty",
		Sound: "Meow!! Meow!!",
	}

	cow := Cow{
		Name:  "Meg",
		Sound: "Moo!! Moo!!",
	}

	animals := []Animal{dog, cat, cow}
	ListAnimal(animals)
}
