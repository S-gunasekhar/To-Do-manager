package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//define the menu items

const (
	ADD_TODO    = 1
	VIEW_TODO   = 2
	SEARCH_TODO = 3
	UPDATE_TODO = 4
	DELETE_TODO = 5
	EXIT        = 6
)

var ToDoList []string
var reader = bufio.NewReader(os.Stdin)

func displayMenu() {

	fmt.Println("1 : To add a new item")
	fmt.Println("2 : To view all the items")
	fmt.Println("3 : Search an items")
	fmt.Println("4 : Update an exiting items")
	fmt.Println("5 : Delete an items")
	fmt.Println("6 : EXIT")
	fmt.Print("\nPlease enter your choice : ")
}

func main() {
	// var choice int

	fmt.Printf("\nWELCOME TO-DO List Manager!!!\n")

	for {
		displayMenu()
		// fmt.Println(choice)
		// fmt.Scanf("%d", &choice)
		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))
		// fmt.Println(input)

		switch choice {
		case ADD_TODO:
			fmt.Println("Add To Do selected")
			AddItem()
		case VIEW_TODO:
			fmt.Println("View To Do selected")
		case SEARCH_TODO:
		case UPDATE_TODO:
		case DELETE_TODO:
		case EXIT:
			fmt.Println("Exiting... Goodbye!")
			return
			// default:
			// 	fmt.Println("Please enter the value between 1 to 6")
		}
	}
}

func AddItem() {
	var item string
	fmt.Print("Enter the item to be added")
	item, _ = reader.ReadString('\n')
	ToDoList = append(ToDoList, item)
	fmt.Println("\nItem Added")
	fmt.Println(ToDoList)
}
