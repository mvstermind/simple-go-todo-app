package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-app/validateinput"
)

func main() {
	fmt.Printf(`
TODO LIST
+------------------------+
|1. Show TODO List       |
|2, Add thing Fo Do      |
|3. Remove From List     |
|4. Exit                 |
+------------------------+

Type a number associated with
the thing that u wanna to: `)
	b := bufio.NewReader(os.Stdin)

	string, err := b.ReadString('\n')
	string = strings.TrimSpace(string)

	if err != nil {
		err := fmt.Errorf("Couldn't read input")
		fmt.Println(err.Error())

	} else {
		validateinput.IsValidInput(string)
	}
}
