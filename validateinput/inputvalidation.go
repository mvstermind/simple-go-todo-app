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

	case text == "3":

	case text == "4":

	default:
		err := fmt.Errorf("Invalid Input")
		fmt.Println(err.Error())
	}

}
