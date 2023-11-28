package main

import (
	"fmt"
)

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
		cmd, err := getCommand()
		if err != nil {
			panic(err) // could be log entry instead of panic
		}

		if err := handleCommand(cmd); err != nil {
			panic(err) // could be log entry instead of panic
		}
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

func getCommand() (string, error) {
	fmt.Printf("\nEnter command > ")

	var cmd string
	if _, err := fmt.Scan(&cmd); err != nil {
		return "", fmt.Errorf("getCommand: scan error: %w", err)
	}

	return cmd, nil
}

func handleCommand(cmd string) error {
	switch cmd {
	case cmdHelp:
		handleHelpCmd()
	case cmdList:
		handleListCmd()
	case cmdCreate:
		return handleCreateCmd()
	case cmdUpdate:
		return handleUpdateCmd()
	case cmdDelete:
		return handleDeleteCmd()
	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		handleHelpCmd()
	}

	return nil
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

func handleCreateCmd() error {
	fmt.Println("Please provide info about new user:")

	fmt.Print("ID: ")
	var id int
	if _, err := fmt.Scan(&id); err != nil {
		return fmt.Errorf("handleCreateCmd: scan error: %w", err)
	}

	fmt.Print("Name: ")
	var name string
	if _, err := fmt.Scan(&name); err != nil {
		return fmt.Errorf("handleCreateCmd: scan error: %w", err)
	}

	_, exists := users[id]
	if exists {
		fmt.Printf("User with ID %d is already exist\n", id)

		return nil
	}

	users[id] = name

	fmt.Println("User successfully created!")

	return nil
}

func handleUpdateCmd() error {
	fmt.Print("Enter user ID to update: ")
	var id int
	if _, err := fmt.Scan(&id); err != nil {
		return fmt.Errorf("handleUpdateCmd: scan error: %w", err)
	}

	name, exists := users[id]
	if !exists {
		fmt.Printf("User with ID %d does not exist\n", id)

		return nil
	}

	fmt.Printf("Current Name: %s\nEnter new name: ", name)
	if _, err := fmt.Scan(&name); err != nil {
		return fmt.Errorf("handleUpdateCmd: scan error: %w", err)
	}

	users[id] = name

	fmt.Println("User successfully updated!")

	return nil
}

func handleDeleteCmd() error {
	fmt.Print("Enter user ID to delete: ")
	var id int
	if _, err := fmt.Scan(&id); err != nil {
		return fmt.Errorf("handleDeleteCmd: scan error: %w", err)
	}

	_, exists := users[id]
	if !exists {
		fmt.Printf("User with ID %d does not exist\n", id)

		return nil
	}

	delete(users, id)

	fmt.Println("User successfully deleted!")

	return nil
}
