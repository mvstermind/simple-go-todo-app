package readdata

import (
	"fmt"
	"os"
)

func DeleteAll() {

	defer func() {
		fmt.Print("\033[H\033[2J") // that shit clears terminal and it's nuts
	}()

	err := os.WriteFile(FilePath, []byte(""), 0666)
	if err != nil {
		err := fmt.Errorf("Couldnt write to file")
		fmt.Println(err.Error())
	}

}
