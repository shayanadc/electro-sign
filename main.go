package main

import (
	"fmt"
	"strconv"
)

type Grid struct {
	pixels [6][36]bool
}

func (g *Grid) TunrOnPixels(input string) {
	left := 0
	right := 1
	input = input + "/"

	for right < len(input) {
		// When we find a letter, process the previous segment
		if input[right] >= 'A' && input[right] <= 'F' || input[right] == '/' {
			// Get row number from left pointer (A=0, B=1, etc.)
			rowNum := int(input[left] - 'A')
			// Get column number from the digits between left+1 and right
			colNum, _ := strconv.Atoi(input[left+1 : right])

			// Set the pixel in the grid
			g.pixels[rowNum][colNum] = true

			// Move left pointer to where right is (the new letter)
			left = right
		}
		// Move right pointer one step forward to start looking for next letter
		right++
	}
}

func (g *Grid) Display() {
	fmt.Println("Displaying the grid:")
	for i := 0; i < 6; i++ {
		for j := 0; j < 36; j++ {
			if g.pixels[i][j] {
				print("*")
			} else {
				print(" ")
			}
		}
		println() // new line after each row
	}
}

func main() {
	sign := &Grid{}
	sign.TunrOnPixels("B1B4B7B8B9B10B13B19B25B26B27B28C1C4C7C13C19C25C28D1D2D3D4D7D8D9D13D19D25D28E1E4E7E13E19E25E28F1F4F7F8F9F10F13F14F15F16F19F20F21F22F25F26F27F28")
	sign.Display()
}
