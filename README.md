# ITEA Golang

Code from lectures of Golang course in IT Education Academy.

## Topics

1. [Hello World](hello_world.go)
2. [Variables](variables.go)
3. [Constants](constants.go)
4. [Enum & iota](enum_iota.go)
5. [If-else](if_else.go)
6. [Switch](switch.go)
7. [For](for.go)
8. [Arrays](arrays.go)
9. [Slices](slices.go)
10. [Map](map.go)
11. [Range](range.go)
12. [Functions](functions.go)
13. [Closures](closures.go)
14. [Pointers](pointers.go)
15. [Structs](structs.go)
16. [Methods](methods.go)
17. [Embedding](embedding.go)
18. [Interfaces](interfaces.go)
19. [Env variables](env_variables.go)
20. [Errors](errors.go)
21. [Defer](defer.go)
22. [Panic, recover](panic_recover.go)
23. [Packages](packages)
24. [Libs](libs)
25. [Database](database)
26. [Unit tests](unit_tests)
27. [Integration tests (godog)](integration_tests)
28. [Type conversion](type_conversion.go)
29. [Type assertion](type_assertion.go)
30. [JSON](json)
31. [HTTP Client](http_client)
32. [HTTP Server](server.go)
33. [Complex HTTP Server](http_server)
34. [Proto](proto)
35. [gRPC](grpc)
36. [Sleep](sleep.go)
37. [Goroutines](goroutines.go)
38. [Wait Group](wait_group.go)
39. [Race condition](race_condition.go)
40. [Atomic](atomic.go)
41. [Mutex](mutex.go)
42. [Channels](channels.go)

## Solutions

Here you can find solutions for part of home works.

4. [Users DB](solutions/4_users_db/main.go)
5. [Student grades](solutions/5_grades/main.go)
5. [Student testing](solutions/5_student_testing/main.go)
7. [Student testing](solutions/7_student_testing)
8. [Integration tests](solutions/8_integration_tests)
8. [Unit tests: student grades](solutions/8_grades_unit_tests)
8. [Unit tests: order](solutions/8_order_unit_tests)
9. [JSON](solutions/9_json)

## Commands

`go build main.go` - compile `main.go` file in current directory

`go run main.go` - compile and run `main.go` file in current directory

`go run .` - compile and run all source files in current directory (must be run in directory with `main()` function)

`go mod init project_name` - create new project with packages support

`go get module_name` - install external library

`go mod tidy` - remove unused libs from `go.mod` & remove `// indirect comment`

`go test -v ./...` - run all tests in the project
