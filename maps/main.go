package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "1",
		"blue":  "2",
		"green": "3",
	}

	printColor(colors)
	delete(colors, "red")
	printColor(colors)
}

func printColor(c map[string]string) {
	for color, num := range c {
		fmt.Println("Color is ", color, "num is", num)
	}
}
