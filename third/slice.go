package slicemain //Change It To main

import "fmt"

func main() {
	var console []string

	console = append(console, "Playstation")
	console = append(console, "PC")
	console = append(console, "Xbox")
	console = append(console, "Sega")
	console = append(console, "Nintendo")

	for _, gaming := range console {
		fmt.Println(gaming)
	}
}
