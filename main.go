package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var A []string

func main() {
	fmt.Println("Welcome to your TODO list !")
	number := 10
	for number != 0 {
		TodoStart()
		reader := bufio.NewReader(os.Stdin)
		choice, er := reader.ReadString('\n')
		if nil != er {
			panic(er)
		}
		choice = strings.TrimSpace(choice)
		fmt.Println("List space: ", number)
		if choice == "1" {
			println("ok")
			number = 0
		}
		if choice == "2" {
			DisplayMyTodoList()
		}
		if choice == "3" {
			fmt.Println("Ah, sure. Like you already didn't have enough. \n" +
				"Write an Item")
			answer, _ := reader.ReadString('\n')
			AddToList(strings.TrimSuffix(answer, "\n"))
		}
		number--

	}
}

func TodoStart() {
	fmt.Println("Select one of the following options:")
	fmt.Println("1. Kill yourself")
	fmt.Println("2. Show your current TODO list")
	fmt.Println("3. Add items to your TODO list")
}

func DisplayMyTodoList() {
	i := 0
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
