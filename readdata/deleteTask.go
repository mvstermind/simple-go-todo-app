package readdata

import (
	"bufio"
	"fmt"
	"os"
)

func DeleteTask() {
	data, err := os.ReadFile(FilePath)
	if err != nil {
		fmt.Printf("Couldn't read file content: %v\n", err)
		return
	}
	fmt.Printf("%v\n", string(data))

	fmt.Print("Which task you want to remove: ")

	b := bufio.NewReader(os.Stdin)

	_, err = b.ReadString('\n')
	if err != nil {
		err := fmt.Errorf("Couldn't read input")
		fmt.Printf(err.Error())
	}
	file, err := os.OpenFile(FilePath, os.O_RDWR|os.O_APPEND, 0666)
	scanner := bufio.NewScanner(file)

	var linesCount int
	for scanner.Scan() {
		linesCount++
	}
	fmt.Printf("%d", linesCount)

}
