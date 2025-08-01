package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var StorageUser []User

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

func runCommand(command string) {
	switch {
	case command == "":
		fmt.Println("Please Enter youre command :")
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

	fmt.Println("please Enter youre password: ")
	scanner.Scan()
	password = scanner.Text()

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
	var ID, email, password string

	fmt.Print("Enter user email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please Enter youre password:")
	scanner.Scan()
	password = scanner.Text()

	notFound := true
	for _, user := range StorageUser {
		if user.Email == email {
			if user.Password == password {
				notFound = true
				fmt.Println("youre are login :")
			} else {
				fmt.Println("the pass or email in incorrect")
			}
		}
	}
	if notFound {
		fmt.Println("the email or password is not corrent ")
		return
	}

	fmt.Print("Enter user password: ")
	scanner.Scan()
	fmt.Println("welcome", ID, email, "youre password is:", password)
}
