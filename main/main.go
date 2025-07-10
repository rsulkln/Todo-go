package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	command := flag.String("command", "no command input", "plaese enter youre command")
	flag.Parse()
	runCommand(*command)
	fmt.Println("StorageUser:", StorageUser)
	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter command: ")
		scanner.Scan()
		*command = scanner.Text()
	}
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var StorageUser []User

func runCommand(command string) {
	switch {
	case command == "create-task":
		createTask()
	case command == "create-category":
		createCategory()
	case command == "create-user":
		createUser()
	case command == "login":
		login()
	case command == "exit":
		os.Exit(0)
	default:
		fmt.Println("No valid command provided. Please use --command to specify a valid command.")
	}
}

func createTask() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Creating a new task...")
	var name, category, duedate string
	fmt.Print("Enter task name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Enter task category: ")
	scanner.Scan()
	category = scanner.Text()

	fmt.Print("Enter task duedate: ")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("youre task is:", name, category, duedate)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("plase Enter youre category youre self: ")
	var category, color string
	fmt.Print("Enter category name: ")
	scanner.Scan()
	category = scanner.Text()

	fmt.Print("Enter category color: ")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("youre category is:", category, color)
}

func createUser() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please Enter youre user youre self: ")
	var name, id, password string
	fmt.Print("Enter user name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("Enter user emai;: ")
	scanner.Scan()
	password = scanner.Text()
	id = password
	fmt.Println("youre user is:", name, password, "and youre ID is:", id)

	user := User{
		ID:       len(StorageUser) + 1,
		Name:     name,
		Password: password,
	}
	StorageUser = append(StorageUser, user)
}
func login() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please Enter youre info for login: ")
	var ID, name, password string
	fmt.Print("Enter user ID: ")
	scanner.Scan()
	ID = scanner.Text()
	fmt.Print("Enter user email: ")
	scanner.Scan()
	name = scanner.Text()
	fmt.Print("Enter user password: ")
	scanner.Scan()
	fmt.Println("welcome", ID, name, "youre password is:", password)
}
