package readdata

import (
	"bufio"
	"fmt"
	"os"
)

var linesCount int = 1

func ReadData() {

	file, err := os.OpenFile(FilePath, os.O_RDONLY, 0666)
	if err != nil {
		err := fmt.Errorf("Error: %v\n", err)
		fmt.Printf(err.Error())
	}
	defer file.Close()

	fmt.Println("Things to do")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		output := fmt.Sprintf("%d. %s", linesCount, string(scanner.Text()))
		fmt.Println(output)
		linesCount++
	}

}
