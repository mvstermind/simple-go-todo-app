package readdata

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const FilePath = "DATA/todo.txt"

var linesCount int = 1

func AddData() {
	fmt.Printf("Add task TODO: ")
	string := readInput()

	file, err := os.OpenFile(FilePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		err := fmt.Errorf("Error: %v\n", err)
		fmt.Printf(err.Error())
	}
	defer file.Close()

	linesCount := getNumLines(file)
	combinedString := fmt.Sprintf("%d. %s", linesCount, string)
	_, err = file.WriteString(combinedString)
	if err != nil {
		err := fmt.Errorf("Error: %v\n", err)
		fmt.Printf(err.Error())
	}

	fmt.Println("Task added!")
	time.Sleep(time.Second)

	defer func() {
		fmt.Print("\033[H\033[2J") // that shit clears terminal and it's nuts
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesCount++
	}
	fmt.Printf("%d", linesCount)

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

func getNumLines(file *os.File) int {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linesCount++
	}
	return linesCount

}
