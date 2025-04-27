package main

import "strconv"

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
