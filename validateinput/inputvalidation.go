package validateinput

import (
	"fmt"
)

func IsValidInput(text string) string {

	switch {
	case text == "1":
		return "1"
	case text == "2":
		return "2"
	case text == "3":
		return "3"
	case text == "4":
		return "4"
	default:
		err := fmt.Errorf("Invalid Input")
		fmt.Println(err.Error())
	}

	return text
}
