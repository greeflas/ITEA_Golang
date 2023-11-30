package main

import "fmt"

type Formatter interface {
	Format(data string) string
}

type JSONFormatter struct{}

func (J JSONFormatter) Format(data string) string {
	return "{" + data + "}"
}

type HTMLFormatter struct{}

func (H HTMLFormatter) Format(data string) string {
	return "<div>" + data + "</div>"
}

func main() {
	// part 1
	var formatter Formatter

	formatter = JSONFormatter{}

	jsonFormatter := formatter.(JSONFormatter)
	fmt.Println(jsonFormatter.Format("45"))

	formatter = HTMLFormatter{}
	jsonFormatter2, ok := formatter.(JSONFormatter)
	if !ok {
		fmt.Println("cannot perform type assertion")
	} else {
		fmt.Println(jsonFormatter2.Format("45"))
	}

	// part 2
	//var x any
	var x interface{}

	x = 2
	fmt.Println(x)

	x = "str"
	fmt.Println(x)

	s := x.(string)
	fmt.Println(s)

	switch val := x.(type) {
	case int:
		fmt.Println("x has type int:", val)
	case string:
		fmt.Println("x has type string:", val)
	case bool:
		fmt.Println("x has type bool:", val)
	default:
		fmt.Println("x type is unknown")
	}
}
