package main

import (
	"fmt"
	"github.com/greeflas/itea_lessons/pb"
	"google.golang.org/protobuf/proto"
	"os"
)

func main() {
	p := &pb.Person{
		Name:  "John Doe",
		Id:    123,
		Email: "john@example.com",
	}

	data, err := proto.Marshal(p)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("person.txt", data, 0644); err != nil {
		panic(err)
	}

	john := new(pb.Person)

	fileContent, err := os.ReadFile("person.txt")
	if err != nil {
		panic(err)
	}

	if err := proto.Unmarshal(fileContent, john); err != nil {
		panic(err)
	}

	fmt.Println(john)
}
