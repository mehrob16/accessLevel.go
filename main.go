package main

import (
	"awesomeProject1/accessLevel"
	"fmt"
)

func main() {
	println("Введите ваш ID:")
	var humanID int
	fmt.Scan(&humanID)

	println("Введите этаж:")
	var availableFloor int
	fmt.Scan(&availableFloor)

	accessLevel.AccessLevel(humanID, availableFloor)

}
