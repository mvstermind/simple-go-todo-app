package readdata

import (
	"fmt"
	"os"
)

func ReadData() {

	data, err := os.ReadFile("DATA/todo.txt")
	if err != nil {
		fmt.Printf("Couldn't read file content: %v\n", err)
		return
	}
	fmt.Printf(string(data))
}

