package readdata

import (
	"fmt"
	"os"
)

func DeleteAll() {

	err := os.WriteFile(FilePath, []byte(""), 0666)
	if err != nil {
		err := fmt.Errorf("Couldnt write to file")
		fmt.Println(err.Error())
	}
}
