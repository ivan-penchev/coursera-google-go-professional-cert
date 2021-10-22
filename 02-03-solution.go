package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, sound string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	m := map[string]*Animal{
		"cow":   &Animal{"grass", "walk", "moo"},
		"bird":  &Animal{"worms", "fly", "peep"},
		"snake": &Animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Print(">")
		animal := ""
		command := ""
		fmt.Scan(&animal, &command)

		if command == "eat" {
			m[animal].Eat()
		} else if command == "move" {
			m[animal].Move()
		} else if command == "speak" {
			m[animal].Speak()
		}
	}
}
