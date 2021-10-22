package main

import (
	"bufio"
	"fmt"
	"go/types"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow types.Nil

func (t Cow) Eat() {
	fmt.Println("grass")
}

func (t Cow) Move() {
	fmt.Println("walk")
}

func (t Cow) Speak() {
	fmt.Println("moo")
}

type Bird types.Nil

func (t Bird) Eat() {
	fmt.Println("worms")
}

func (t Bird) Move() {
	fmt.Println("fly")
}

func (t Bird) Speak() {
	fmt.Println("peep")
}

type Snake types.Nil

func (t Snake) Eat() {
	fmt.Println("mice")
}
func (t Snake) Move() {
	fmt.Println("slither")
}
func (t Snake) Speak() {
	fmt.Println("hsss")
}

func getUserInput() []string {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	return strings.Fields(text)
}

func getAnimalType(animalType string) Animal {
	switch animalType {
	case "cow":
		return Cow{}
	case "bird":
		return Bird{}
	case "snake":
		return Snake{}
	default:
		return nil
	}
}

func main() {
	animals := make(map[string]Animal)
	fmt.Println("You can create animals or query information about the animals you have created before")
	fmt.Println("Example: <newanimal> <animal_name> [cow|bird|snake]")
	fmt.Println("         <query> <animal_name> [eat|move|speak]")
	for {
		request := getUserInput()
		if len(request) != 3 {
			fmt.Println("Unknown request")
			continue
		}
		switch request[0] {
		case "query":
			// <name_of_animal> <request_information>
			animal := animals[request[1]]
			if animal == nil {
				fmt.Println("Animal not exist: ", animals[request[1]])
				continue
			}
			switch request[2] {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				{
					fmt.Println("Requested information is not supported: ", request[2])
					continue
				}
			}
		case "newanimal":
			// <newanimal> <name_of_animal> <animal_type>
			animal := animals[request[1]]
			if animal != nil {
				fmt.Println("Requested animal name exist: ", request[1])
				continue
			}
			animals[request[1]] = getAnimalType(request[2])
			fmt.Println("Created it!")

		default:
			fmt.Println("Not supported request: ", request)
		}
	}
}
