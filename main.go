package main

import (
	"bufio"
	"fmt"
	"github.com/MasterDimmy/go-cls"
	"os"
	"strings"
)

var A []string

func main() {
	fmt.Println("Welcome to your TODO list !")
	choices := []string{"1", "2", "3", "4"}
	number := 10
	for number != 0 {
		TodoStart()
		//starting with the couple of begging print statements.
		reader := bufio.NewReader(os.Stdin)
		choice, er := reader.ReadString('\n')
		if nil != er {
			panic(er)
		}
		//checking to see if a choice is valid
		validChoice := false
		for i := 0; i < len(choices); i++ {
			if strings.TrimSpace(choice) == choices[i] {
				validChoice = true
				break
			}

		}
		if !validChoice {
			fmt.Println("Invalid Choice. Try Again")
			continue
		}
		//If 1, break the loop
		choice = strings.TrimSpace(choice)
		fmt.Println("List space: ", number)
		if choice == "1" {
			println("ok")
			break
		}
		//Displaying all items from the todo list if 2
		if choice == "2" {
			DisplayMyTodoList()
		}
		//if 3, add the item to a TODO list
		if choice == "3" {
			fmt.Println("Ah, sure. Like you already didn't have enough. \n" +
				"Write an Item")
			answer, _ := reader.ReadString('\n')
			AddToList(strings.TrimSuffix(answer, "\n"))
		}
		//If choice is 4, clear the console.
		if choice == "4" {
			ClearConsole()
		}
		number--

	}
}

func TodoStart() {
	fmt.Println("Select one of the following options:")
	fmt.Println("1. Kill yourself")
	fmt.Println("2. Show your current TODO list")
	fmt.Println("3. Add items to your TODO list")
	fmt.Println("4. Clear the Console.")
}

func DisplayMyTodoList() {
	i := 1
	for _, item := range A {
		if item != "" {
			fmt.Println(i, ": ", item)
			i++
		}
	}
}

func AddToList(newItem string) {
	A = append(A, newItem)
}
func ClearConsole() {
	cls.CLS()
}
