package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["Red"] = "#FF0000"
	colors["White"] = "#FFFFFF"
	colors["Black"] = "#000000"

	otherColors := map[string]string{
		"Yellow": "#aaaaaa",
		"Blue":   "#bbbbbbb",
	}
	printColors(colors)
	printColors(otherColors)

}

func printColors(colors map[string]string) {
	for color, value := range colors {
		fmt.Println("Color code of ", color, " is ", value)
	}
}
