package main

import "fmt"

func main() {

	luffy, _ := GetPirate("luffy")
	law, _ := GetPirate("law")
	roger, _ := GetPirate("roger")

	printDetails(luffy)
	printDetails(law)
	printDetails(roger)

}

func printDetails(p IPirate) {
	fmt.Printf("Pirate: %s", p.getName())
	fmt.Println()
	fmt.Printf("Power: %s", p.getPower())
	fmt.Println()
}
