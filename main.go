package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-app/validateinput"
)

// clear list has to be like this cuz otherwise it will be offset, dunno why this happens
func main() {

	fmt.Printf(`
TODO LIST
+------------------------+
|1. Show TODO List       |
|2. Add thing Fo Do      |
|3. Remove From List     |
|4. Clear List		 | 
|5. Exit                 |
+------------------------+
Chose action: `)
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
