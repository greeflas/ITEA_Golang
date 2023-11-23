package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println(id.String())

	parsedId, err := uuid.Parse("4e4adb3a-4ba4-44f1-a778-c1263349754d")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(parsedId.String())

	_, err = uuid.Parse("invalid id")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(uuid.MustParse("69e8eafc-a72d-4abe-addd-e421ef9f659c").String())

	uuid.MustParse("invalid id")
}
