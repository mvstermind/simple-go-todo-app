package readdata

import (
	"bufio"
	"fmt"
	"os"
)

func AddData() {
	fmt.Printf("Add task TODO: ")
	string := readInput()

	file, err := os.OpenFile("DATA/todo.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		err := fmt.Errorf("Error: %v\n", err)
		fmt.Printf(err.Error())
	}
	defer file.Close()

	_, err = file.WriteString(string)
	if err != nil {
		err := fmt.Errorf("Error: %v\n", err)
		fmt.Printf(err.Error())
	}

	fmt.Println("Task added!")
}

func readInput() string {

	b := bufio.NewReader(os.Stdin)

	string, err := b.ReadString('\n')
	if err != nil {
		err := fmt.Errorf("Couldn't read input")
		fmt.Printf(err.Error())
	}
	return string
}
