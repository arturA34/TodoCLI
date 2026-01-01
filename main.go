package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	title       string
	description string
	status      string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var choice int
	var todos []Todo
	for {
		fmt.Print(" /████████ /██████  /███████   /██████         /██████  /██       /██████\n|__  ██__//██__  ██| ██__  ██ /██__  ██       /██__  ██| ██      |_  ██_/\n   | ██  | ██  \\ ██| ██  \\ ██| ██  \\ ██      | ██  \\__/| ██        | ██  \n   | ██  | ██  | ██| ██  | ██| ██  | ██      | ██      | ██        | ██  \n   | ██  | ██  | ██| ██  | ██| ██  | ██      | ██      | ██        | ██  \n   | ██  | ██  | ██| ██  | ██| ██  | ██      | ██    ██| ██        | ██  \n   | ██  |  ██████/| ███████/|  ██████/      |  ██████/| ████████ /██████\n   |__/   \\______/ |_______/  \\______/        \\______/ |________/|______/ ")
		fmt.Print("\n\n" +
			"1. Create Todo\n" +
			"2. List All Todos\n" +
			"3. Update Todo\n" +
			"4. Delete Todo\n" +
			"5. Change Todo Status\n\n" +
			"What do you want to do : ")
		choice, _ = strconv.Atoi(readLine(reader))
		switch choice {
		case 1:
			var todoName, todoDescription string

			fmt.Print("Enter the name of the todo: ")
			todoName = readLine(reader)
			fmt.Print("Enter the description of the todo (optional): ")
			todoDescription = readLine(reader)
			todos = createTodo(todoName, todoDescription, todos)

			todos = createAnotherTodo(reader, todos)
		case 2:
			listAllTodos(todos)
			fmt.Print("\nActions:\n" +
				"1. Return to the menu\n" +
				"What do you want to do next: ")
			choice, _ = strconv.Atoi(readLine(reader))
			switch choice {
			case 1:
				break
			}
		case 3:
			var name, description string
			var id int
			if listAllTodos(todos) {
				fmt.Print("\nEnter the todo id that you want to update: ")
				id, _ = strconv.Atoi(readLine(reader))
				fmt.Print("Enter the new name of the todo: ")
				name = readLine(reader)
				fmt.Print("Enter the new description of the todo (optional): ")
				description = readLine(reader)
				updateTodo(name, description, id, todos)
			}
		case 4:
			var id int
			if listAllTodos(todos) {
				fmt.Print("\nEnter the todo id that you want to delete: ")
				id, _ = strconv.Atoi(readLine(reader))
				todos = deleteTodo(id, todos)
			}
		case 5:
			if listAllTodos(todos) {
				fmt.Print("\nEnter the todo ID whose status you want to change: ")
				id, _ := strconv.Atoi(readLine(reader))
				fmt.Print("Enter a new todo status: ")
				status := readLine(reader)
				todos[id].status = status
			}
		}
	}
}

func createTodo(todoName string, todoDescription string, todos []Todo) []Todo {
	if todoName == "" {
		return todos
	}
	todoModel := Todo{todoName, todoDescription, "new"}
	return append(todos, todoModel)
}

func updateTodo(todoName string, todoDescription string, todoId int, todos []Todo) {
	if todoName != "" {
		todos[todoId].title = todoName
	}
	if todoDescription != "" {
		todos[todoId].description = todoDescription
	}
}

func deleteTodo(todoId int, todos []Todo) []Todo {
	return append((todos)[:todoId], (todos)[todoId+1:]...)
}

func readLine(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

func listAllTodos(todos []Todo) bool {
	if len(todos) == 0 {
		fmt.Println("There are no todos, returning to main menu")
		return false
	}
	fmt.Println("+------+----------+--------------------+--------+")
	fmt.Println("|  ID  |   Name   |     Description    | Status |")
	fmt.Println("+------+----------+--------------------+--------+")
	for i, todo := range todos {
		fmt.Printf("|%-6d|%-10s|%-20s|%-8s|\n", i, todo.title, todo.description, todo.status)
		fmt.Println("+------+----------+--------------------+--------+")
	}
	return true
}

func createAnotherTodo(reader *bufio.Reader, todos []Todo) []Todo {
	var todoName, todoDescription string
	var choice int
	fmt.Print("1. Return to the menu\n" +
		"2. Create another todo\n" +
		"What do you want to do next: ")
	choice, _ = strconv.Atoi(readLine(reader))
	switch choice {
	case 1:
		return todos
	case 2:
		fmt.Print("Write the name of the todo: ")
		todoName = readLine(reader)
		fmt.Print("Write the description of the todo (optional): ")
		todoDescription = readLine(reader)
		todos = createTodo(todoName, todoDescription, todos)
		return createAnotherTodo(reader, todos)
	default:
		return todos
	}
}
