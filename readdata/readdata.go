package readdata

import (
	"fmt"
	"os"
)

func ReadData() {

	data, err := os.ReadFile(FilePath)
	if err != nil {
		fmt.Printf("Couldn't read file content: %v\n", err)
		return
	}
	fmt.Printf(string(data))
}
