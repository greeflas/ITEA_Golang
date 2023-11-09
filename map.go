package main

import "fmt"

func main() {
	users := map[int]string{
		100: "John",
		101: "Jane",
	}

	fmt.Println(users)
	fmt.Println("Len:", len(users))

	fmt.Println(users[101])

	users[102] = "Jack"
	fmt.Println(users)

	delete(users, 100)
	fmt.Println(users)

	name, exists := users[103]
	fmt.Println(name, exists)

	name, exists = users[102]
	fmt.Println(name, exists)
}
