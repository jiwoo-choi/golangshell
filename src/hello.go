package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// error handling
//https://kimmachinegun.github.io/2018-08-29/error-%EB%A9%8B%EC%A7%80%EA%B2%8C-%EB%8B%A4%EB%A3%A8%EA%B8%B01/

//struct
//https://gobyexample.com/structs

//interface
//https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c

//fmt what is %q.
//https://golang.org/pkg/fmt/

//https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go
// http://golang.site/go/article/18-Go-%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4
//https://mingrammer.com/gobyexample/interfaces/

type command struct {
	mainCommand            string
	contents               string
	nextCommandAvailiblity bool
	nextCommandLine        *command
}

// 다형성지원?

func main() {

	for {

		fmt.Printf("jiwoo-shell > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command, _ := parser(scanner.Text())

		fmt.Printf(command.mainCommand)

		/*for scanner.Scan() {
					fmt.Println(scanner.Text())
		    }*/

		/*
			var s1 string
			n, _ := fmt.Scan(&s1)
			if strings.ToLower(s1) == "exit" {
				break
			}
			fmt.Println("입력 개수:", n)
		*/
	}

}

//https://gobyexample.com/structs
//struct intiliziation.

func parser(originalCommand string) (*command, error) {
	splittedCommand := strings.Split(originalCommand, " ")
	commandStruct := command{mainCommand: splittedCommand[0], contents: splittedCommand[1]}
	if strings.Contains(originalCommand, "&") || strings.Contains(originalCommand, "|") {
		commandStruct.nextCommandAvailiblity = true
	}
	//& or something..?
	return &commandStruct, nil
}
