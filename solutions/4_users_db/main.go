package main

import "fmt"

const (
	cmdHelp   = "help"
	cmdList   = "list"
	cmdCreate = "create"
	cmdUpdate = "update"
	cmdDelete = "delete"
)

var cmdDescriptions = map[string]string{
	cmdHelp:   "this help",
	cmdList:   "show all users in database",
	cmdCreate: "create new user and add to database",
	cmdUpdate: "update user data in database",
	cmdDelete: "delete user from database",
}

var users = make(map[int]string)

func main() {
	printWelcome()

	for {
		cmd := getCommand()
		handleCommand(cmd)
	}
}

func printWelcome() {
	const msg = "Welcome to Users Database v1"
	borderLen := len(msg) + 2

	printBorder(borderLen)
	fmt.Printf(" %s \n", msg)
	printBorder(borderLen)

	fmt.Printf("\nType 'help' to get list of available commands.\n")
}

func printBorder(len int) {
	for i := 0; i < len; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}

func getCommand() string {
	fmt.Printf("\nEnter command > ")

	var cmd string
	fmt.Scan(&cmd)

	return cmd
}

func handleCommand(cmd string) {
	switch cmd {
	case cmdHelp:
		handleHelpCmd()
	case cmdList:
		handleListCmd()
	case cmdCreate:
		handleCreateCmd()
	case cmdUpdate:
		handleUpdateCmd()
	case cmdDelete:
		handleDeleteCmd()
	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		handleHelpCmd()
	}
}

func handleHelpCmd() {
	fmt.Println("Available commands:")

	for cmd, description := range cmdDescriptions {
		fmt.Printf("%s - %s.\n", cmd, description)
	}
}

func handleListCmd() {
	if len(users) == 0 {
		fmt.Println("Database is empty")

		return
	}

	fmt.Println("Users in database:")

	for id, name := range users {
		fmt.Printf("\nID: %d\nName: %s\n", id, name)
	}
}

func handleCreateCmd() {
	fmt.Println("Please provide info about new user:")

	fmt.Print("ID: ")
	var id int
	fmt.Scan(&id)

	fmt.Print("Name: ")
	var name string
	fmt.Scan(&name)

	_, exists := users[id]
	if exists {
		fmt.Printf("User with ID %d is already exist\n", id)

		return
	}

	users[id] = name

	fmt.Println("User successfully created!")
}

func handleUpdateCmd() {
	fmt.Print("Enter user ID to update: ")
	var id int
	fmt.Scan(&id)

	name, exists := users[id]
	if !exists {
		fmt.Printf("User with ID %d does not exist\n", id)

		return
	}

	fmt.Printf("Current Name: %s\nEnter new name: ", name)
	fmt.Scan(&name)

	users[id] = name

	fmt.Println("User successfully updated!")
}

func handleDeleteCmd() {
	fmt.Print("Enter user ID to delete: ")
	var id int
	fmt.Scan(&id)

	_, exists := users[id]
	if !exists {
		fmt.Printf("User with ID %d does not exist\n", id)

		return
	}

	delete(users, id)

	fmt.Println("User successfully deleted!")
}
