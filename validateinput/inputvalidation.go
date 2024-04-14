package validateinput

import (
	"fmt"
	"todo-app/readdata"
)

func IsValidInput(text string) {

	switch {
	case text == "1":
		readdata.ReadData()

	case text == "2":
		readdata.AddData()

	case text == "3":
		readdata.DeleteTask()

	case text == "4":
		readdata.DeleteAll()

	case text == "5":
	default:
		err := fmt.Errorf("Invalid Input")
		fmt.Println(err.Error())
	}

}
