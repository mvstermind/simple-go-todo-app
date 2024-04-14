package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"todo-app/validateinput"
)

func main() {
	fmt.Printf(`
TODO LIST
+------------------------+
|1. Show TODO List       |
|2. Add thing Fo Do      |
|3. Remove From List     |
|4. Exit                 |
+------------------------+

Type a number associated with
the thing that u wanna to: `)
	r := bufio.NewReader(os.Stdin)

	b, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		validateinput.IsValidInput(b)
	}
}
