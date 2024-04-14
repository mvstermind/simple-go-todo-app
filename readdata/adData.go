package readdata

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

const FilePath = "DATA/todo.txt"

func AddData() {
	fmt.Printf("Add task TODO: ")
	string := readInput()

	file, err := os.OpenFile(FilePath, os.O_RDWR|os.O_APPEND, 0666)
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

	scanner := bufio.NewScanner(io.Reader(file))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("Task added!")
	time.Sleep(time.Second)

	defer func() {
		fmt.Print("\033[H\033[2J") // that shit clears terminal and it's nuts
	}()

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
