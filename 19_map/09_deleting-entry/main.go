package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two") // delete entry from the map
	fmt.Println(myGreeting)
}
