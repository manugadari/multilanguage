package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// calculateRectangleArea calculates the area of a rectangle
func calculateRectangleArea(length, width float64) float64 {
	return length * width
}

// calculateCircleArea calculates the area of a circle
func calculateCircleArea(radius float64) float64 {
	return math.Pi * (radius * radius)
}

func main() {
	fmt.Println("Welcome to the Area Calculator!")

	// Ask the user for the type of shape they want to calculate the area for
	fmt.Print("Enter 'rectangle' or 'circle': ")
	var shapeType string
	fmt.Scanln(&shapeType)

	// Validate the user's input
	for strings.ToLower(shapeType) != "rectangle" && strings.ToLower(shapeType) != "circle" {
		fmt.Print("Invalid input. Please enter 'rectangle' or 'circle': ")
		fmt.Scanln(&shapeType)
	}

	// Ask the user for the dimensions of the shape
	if strings.ToLower(shapeType) == "rectangle" {
		fmt.Print("Enter the length of the rectangle: ")
		var length float64
		_, err := fmt.Scanln(&length)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Print("Enter the width of the rectangle: ")
		var width float64
		_, err = fmt.Scanln(&width)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		area := calculateRectangleArea(length, width)
		fmt.Printf("The area of the rectangle is %.2f square units.\n", area)
	} else if strings.ToLower(shapeType) == "circle" {
		fmt.Print("Enter the radius of the circle: ")
		var radius float64
		_, err := fmt.Scanln(&radius)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		area := calculateCircleArea(radius)
		fmt.Printf("The area of the circle is %.2f square units.\n", area)
	}

	fmt.Println("Thank you for using the Area Calculator!")
}
