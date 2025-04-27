## **Designing an Evolving System: From Pixels to Patterns**

In software development, we constantly walk a tightrope between what the system needs *now* and what it *might* need in the future. Focus too much on the present, and you build something clean—but rigid. Plan too far ahead, and you risk building castles in the sky.

Good design isn’t static. It evolves. It adapts. It listens. It’s less about solving all the future’s problems today and more about making choices that keep doors open tomorrow.

I'll demonstrate this philosophy by building a simple—but extensible—system that simulates an **electronic signboard**.

---

## **The Problem: A Digital Signboard**

We’re designing a console-based application that simulates an electronic sign with the following requirements:

- A grid of **6 rows by 36 columns**, where each cell (pixel) can be either **On** or **Off**
- Users can input a sequence of pixel coordinates (e.g., `"B1B4B7"`) to form a **View**
- The system can **store multiple views** in memory
- It can **display all views** or **clear the memory**

Pixels are identified using a letter-number format: `A0` (top-left) to `F35` (bottom-right).

---

#### Example

The following sequence:

```
B1B4B7B8B9B10B13B19B25B26B27B28C1C4C7C13C19C25C28D1D2D3D4D7D8D9D13D19D25D28E1E4E7E13E19E25E28F1F4F7F8F9F10F13F14F15F16F19F20F21F22F25F26F27F28
```
It should give the following result:

```
*   * ***** *     *      ****
*   * *     *     *     *    *
***** ***** *     *     *    *
*   * *     *     *     *    *
*   * ***** ***** *****  ****
```

## **Step One: Choosing the Right Data Structure**

At the heart of any software system lies its data model—how information flows in, how it’s structured internally, and how it’s exposed.

I opted for a **2D array** of booleans—6 rows by 36 columns.

Why?

- ✅ **Efficiency** – Each boolean uses only 1 byte. Using **strings** or larger types would bloat memory use.
- ✅ **Clarity** – A pixel is either `true` (on) or `false` (off). No ambiguity.
- ✅ **Speed** – Access is direct, looping is clean, and performance is predictable.

This decision keeps the system lean, fast, and easy to reason about while these decisions will be affect where does the system design go.

```Golang
type Grid struct {
	pixels [6][36]bool
}
```

---

## **Core Components: Bringing the Sign to Life**

### `TurnOnPixels` Method

This method takes a string of pixel coordinates (e.g., `"B1B4B7"`) and parses it into action.

- Uses a **two-pointer approach**:
  - Left pointer grabs the row (A–F)
  - Right pointer scans forward for the next letter to grab the column number in between
- Converts these into array indices and sets the corresponding pixels to `true`

It’s lightweight, direct, and avoids unnecessary overhead like regex parsing.

```Golang
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
            colNum, _ := strconv.Atoi(input[left+1:right])
            
            // Set the pixel in the grid
            g.pixels[rowNum][colNum] = true
            
            // Move left pointer to where right is (the new letter)
            left = right
        } 
            // Move right pointer one step forward to start looking for next letter
        right++
    }
}
```
---

### `Display` Method

The `Display` method shows the current state of the sign using **ASCII art**:

- `'*'` for pixels that are **on**
- `' '` (space) for those that are **off**

It loops through the grid and prints each row, giving a console-friendly preview of what the sign looks like.

```Golang
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
```


---

## **Architecture & Design Considerations**

While the current implementation works well for our initial goals, several areas stand out for future refinement:

**Grid Size**: Should we allow dynamic dimensions instead of hardcoding?

**Pixel Representation**: Is boolean sufficient long-term, or will we need richer pixel states?

**Display Strategy**: Can ASCII output scale, or should we explore more flexible rendering interfaces?

These questions open the door to enhancements as the system evolves—and we’ll revisit them when the time is right.

## ***Build Small, Think Big***

What we’ve created so far is just the beginning. But it’s not a throwaway prototype—it’s a **solid core**. Every choice we made was intentional: small, purposeful, and open to change.

In the end, that’s what good software design is all about.

---