package main

import "fmt"

func main() {

	v := 217

	if v < 300 {

		fmt.Println("Smaller")

	} else {

		fmt.Println("Greater")

	}

	//nested statement

	if v < 400 {
		if v < 300 {
			fmt.Println("Smaller")
		}

	}

	// else if

	if v > 250 {

		fmt.Println("Bigger")
	} else if v == 217 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Smaller")
	}

}
