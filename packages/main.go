package main

import (
	"fmt"
	mytime "github.com/greeflas/itea_golang/datetime"
	"github.com/greeflas/itea_golang/filesystem/directory"
	"github.com/greeflas/itea_golang/filesystem/file"
	"github.com/greeflas/itea_golang/first"
	//"github.com/greeflas/itea_golang/first/internal"
	"github.com/greeflas/itea_golang/first/second"
	"time"
)

func main() {
	now := time.Now()

	dt := mytime.New(now)
	fmt.Println(dt)

	dt2 := dt.StartOfHour()
	fmt.Println(dt2)

	dt3 := dt.StartOfDay()
	fmt.Println(dt3)

	//fmt.Println(dt.inner)
	//fmt.Println(dt.startOfTemplate(0, 0, 0, 0))

	fmt.Println()

	initialize()

	dirsCount, err := directory.CountDirs("./test")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Directories count:", dirsCount)

	filesCount, err := file.CountFiles("./test")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Files count:", filesCount)

	fmt.Println()

	first.First()
	second.Second()
	//fmt.Println(internal.SecretMessage())
}
