package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	Duedate    time.Time
	CategoryID int
	IsDone     bool
	UserId     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

const userStoragePath = "user.txt"

var (
	StorageUser        []User
	authenticatedUsers *User
	taskStorage        []Task
	categoryStorage    []Category
	file               *os.File
)

func main() {
	loadUserStorageFromFile()
	command := flag.String("command", "no input or nut found", "command to execute")
	flag.Parse()
	//fmt.Printf("userStorage: %+v\n", StorageUser)

	for {
		runCommand(*command)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another commmand >>>  ")
		scanner.Scan()
		*command = scanner.Text()
	}
}
func (u User) print() {
	fmt.Println("user:", u.Name, "ID:", u.ID, "Email:", u.Email)
}
func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUsers == nil {
		login()
	}
	//if authenticatedUsers == nil {
	//	return
	//}

	switch command {
	case "register-user":
		registerUser()
	//case "login":
	//	login()
	case "create-category":
		createCategory()
	case "create-task":
		createTask()
	case "list-task":
		listTask()
	case "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)
	case "":
		registerUser()
	default:
		fmt.Println("command not found ", command)
		return
	}
}
func registerUser() {
	var name, email, password, password2 string
	var id int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your name: ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Enter your email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Enter your password: ")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("Enter your confirm password: ")
	scanner.Scan()
	password2 = scanner.Text()
	if password != password2 || len(password) < 5 {
		fmt.Println("Passwords don't match or not good pass ")
		return

	} else {
		user := User{
			Name:     name,
			Email:    email,
			ID:       len(StorageUser) + 1,
			Password: password,
		}
		fmt.Println("len user storage", len(StorageUser))
		StorageUser = append(StorageUser, user)

		var oErr error
		file, oErr = os.OpenFile(userStoragePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if oErr != nil {
			fmt.Println("you have an error", oErr)

			return
		}
		data := fmt.Sprintf("id: %d\nname: %s\nemail: %s\npassword: %s\n\n",
			user.ID, user.Name, user.Email, user.Password)
		var b = []byte(data)

		numberOfWrite, wErr := file.Write(b)
		if wErr != nil {
			fmt.Println("you have an error", wErr)
			fmt.Printf("Number of bytes written: %d\n", numberOfWrite)

			return
		}
	}

	fmt.Printf("id: %v welcome: %s youre email is: %s \n\n", id, name, email)
}

func createTask() {
	if authenticatedUsers != nil {
		authenticatedUsers.print()
	}

	var title, category, duedate string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your title: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Enter your due date: ")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("Enter your category id: ")
	scanner.Scan()
	category = scanner.Text()
	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("Invalid categoryid %v\n ", err)

		return
	}
	IsFaund := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUsers.ID {
			IsFaund = true
			break
		}
	}
	if !IsFaund {
		fmt.Println("category not found ")
		return
	}

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		Duedate:    time.Now(),
		CategoryID: categoryID,
		IsDone:     false,
		UserId:     authenticatedUsers.ID,
	}

	//fmt.Println("len:", len(taskStorage))
	taskStorage = append(taskStorage, task)

	fmt.Printf("✅ Task created successfully!\n"+
		"📂 Category: %v\n"+
		"📝 Title: %v\n"+
		"📅 Due Date: %v\n",
		categoryID, title, duedate)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string
	fmt.Print("Enter your categoy  title: ")
	scanner.Scan()
	title = scanner.Text()
	fmt.Print("Enter your category color: ")
	scanner.Scan()
	color = scanner.Text()

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUsers.ID,
	}
	categoryStorage = append(categoryStorage, category)
	fmt.Println(len(categoryStorage))
	fmt.Printf("youre title is %s || color is %s\n", title, color)
}

func login() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Print("Enter your email: ")
	scanner.Scan()
	email = scanner.Text()

	fmt.Print("Enter your password: ")
	scanner.Scan()
	password = scanner.Text()

	var found = false

	for _, user := range StorageUser {
		if user.Email == email && user.Password == password {
			fmt.Printf("Login successful! Welcome %s\n", user.Name)
			found = true
			authenticatedUsers = &user

			break
		}
	}
	if !found {
		fmt.Println("Login failed: email or password is incorrect")

		return
	}
}
func listTask() {
	if authenticatedUsers == nil {
		fmt.Println("Please login first!")
		return
	}

	fmt.Printf("Tasks for user %s:\n", authenticatedUsers.Name)
	fmt.Println("====================")

	found := false
	for _, task := range taskStorage {
		if task.UserId == authenticatedUsers.ID {
			found = true

			status := "Not Done"
			if task.IsDone {
				status = "Done"
			}
			fmt.Printf("Title: %s\nStatus: %s\nCategory ID: %d\nDue Date: %s\n\n",
				task.Title, status, task.CategoryID, task.Duedate.Format("2006-01-02"))
		}
	}

	if !found {
		fmt.Println("No tasks found for this user")
	}
}
func loadUserStorageFromFile() {
	file, err := os.Open(userStoragePath)
	if err != nil {
		fmt.Println("userStoragePath doesn't exist", err)

		return
	}
	var data = make([]byte, 1024)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("userStoragePath doesn't exist", oErr)

		return
	}
	var dataString = string(data)

	//strings.Replace(dataString, "\n", " ", -1)
	sliceString := strings.Split(dataString, "\n")
	for index, u := range sliceString {
		if u == "" {
			continue
		}
		fmt.Println("user Line:", index, "user:", u)
		//StorageUser = append(StorageUser, u)
	}
	//fmt.Println("this is you're data:\n", sliceString)
}
