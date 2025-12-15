package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Image struct {
	title string
	data  [][]rune
}

type Picture struct {
	tiles []*Image
}

type Tile struct {
	id    int
	image *Image
}

func MakePicture(imagedata []string) *Picture {
	images := []*Image{}
	data := []string{}
	for _, line := range imagedata {
		if len(line) == 0 {
			images = append(images, MakeImage(data))
			data = []string{}
		} else {
			data = append(data, line)
		}
	}
	images = append(images, MakeImage(data))
	return &Picture{tiles: images}
}

func MakeImage(image []string) *Image {
	title := image[0]
	data := [][]rune{}
	for i := 1; i < len(image); i++ {
		data = append(data, []rune(image[i]))
	}
	return &Image{title: title, data: data}
}

func (i *Image) GetBorder() [][]rune {
	top := i.data[0]
	bottom := i.data[9]
	left := []rune{}
	right := []rune{}

	for _, row := range i.data {
		left = append(left, row[0])
		right = append(right, row[9])
	}

	return [][]rune{top, right, bottom, left}
}

func reverse(orig []rune) []rune {
	newval := []rune{}
	for i := len(orig) - 1; i >= 0; i-- {
		newval = append(newval, orig[i])
	}
	return newval
}

func (i *Image) HasBorderWith(other *Image) bool {
	ourBorders := i.GetBorder()
	theirBorders := other.GetBorder()
	for _, ourborder := range ourBorders {
		for _, theirborder := range theirBorders {
			if string(ourborder) == string(theirborder) || string(reverse(ourborder)) == string(theirborder) {
				return true
			}
		}
	}
	return false
}

func (p *Picture) FindCorners() []*Image {
	corners := []*Image{}
	for i, origTile := range p.tiles {
		count := 0
		for j, compareTile := range p.tiles {
			if i != j {
				if origTile.HasBorderWith(compareTile) {
					count++
				}
			}
		}
		if count == 2 {
			corners = append(corners, origTile)
		}
	}
	return corners
}

// Rotate image 90 degrees clockwise
func (i *Image) Rotate() *Image {
	size := len(i.data)
	newData := make([][]rune, size)
	for row := 0; row < size; row++ {
		newData[row] = make([]rune, size)
		for col := 0; col < size; col++ {
			newData[row][col] = i.data[size-1-col][row]
		}
	}
	return &Image{title: i.title, data: newData}
}

// Flip image horizontally
func (i *Image) FlipH() *Image {
	size := len(i.data)
	newData := make([][]rune, size)
	for row := 0; row < size; row++ {
		newData[row] = make([]rune, size)
		for col := 0; col < size; col++ {
			newData[row][col] = i.data[row][size-1-col]
		}
	}
	return &Image{title: i.title, data: newData}
}

// Flip image vertically
func (i *Image) FlipV() *Image {
	size := len(i.data)
	newData := make([][]rune, size)
	for row := 0; row < size; row++ {
		newData[row] = make([]rune, size)
		for col := 0; col < size; col++ {
			newData[row][col] = i.data[size-1-row][col]
		}
	}
	return &Image{title: i.title, data: newData}
}

// Get all possible orientations of an image
func (i *Image) GetOrientations() []*Image {
	orientations := []*Image{}
	current := i

	// 4 rotations
	for r := 0; r < 4; r++ {
		orientations = append(orientations, current)
		current = current.Rotate()
	}

	// Flip and 4 more rotations
	current = i.FlipH()
	for r := 0; r < 4; r++ {
		orientations = append(orientations, current)
		current = current.Rotate()
	}

	return orientations
}

// Get specific border (0=top, 1=right, 2=bottom, 3=left)
func (i *Image) GetBorderSide(side int) []rune {
	borders := i.GetBorder()
	return borders[side]
}

// Check if two borders match
func bordersMatch(b1, b2 []rune) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

// Assemble the full image from tiles
func (p *Picture) AssembleImage() [][]rune {
	// Find grid size
	gridSize := int(math.Sqrt(float64(len(p.tiles))))

	// Create grid to hold arranged tiles
	grid := make([][]*Image, gridSize)
	for i := range grid {
		grid[i] = make([]*Image, gridSize)
	}

	used := make(map[string]bool)

	// Find a corner to start with
	corners := p.FindCorners()

	// Try different orientations of the first corner
	for _, orientation := range corners[0].GetOrientations() {
		grid[0][0] = orientation
		used[corners[0].title] = true

		if p.fillGrid(grid, used, 0, 1, gridSize) {
			break
		}

		used[corners[0].title] = false
		grid[0][0] = nil
	}

	// Remove borders and combine tiles
	return p.combineTiles(grid, gridSize)
}

// Recursively fill the grid
func (p *Picture) fillGrid(grid [][]*Image, used map[string]bool, row, col, size int) bool {
	if row == size {
		return true
	}

	nextRow, nextCol := row, col+1
	if nextCol == size {
		nextRow++
		nextCol = 0
	}

	for _, tile := range p.tiles {
		if used[tile.title] {
			continue
		}

		for _, orientation := range tile.GetOrientations() {
			valid := true

			// Check left neighbor
			if col > 0 {
				leftTile := grid[row][col-1]
				if !bordersMatch(leftTile.GetBorderSide(1), orientation.GetBorderSide(3)) {
					valid = false
				}
			}

			// Check top neighbor
			if valid && row > 0 {
				topTile := grid[row-1][col]
				if !bordersMatch(topTile.GetBorderSide(2), orientation.GetBorderSide(0)) {
					valid = false
				}
			}

			if valid {
				grid[row][col] = orientation
				used[tile.title] = true

				if p.fillGrid(grid, used, nextRow, nextCol, size) {
					return true
				}

				used[tile.title] = false
				grid[row][col] = nil
			}
		}
	}

	return false
}

// Combine tiles into single image (removing borders)
func (p *Picture) combineTiles(grid [][]*Image, gridSize int) [][]rune {
	tileSize := len(grid[0][0].data)
	innerSize := tileSize - 2 // Remove borders
	fullSize := gridSize * innerSize

	result := make([][]rune, fullSize)
	for i := range result {
		result[i] = make([]rune, fullSize)
	}

	for gridRow := 0; gridRow < gridSize; gridRow++ {
		for gridCol := 0; gridCol < gridSize; gridCol++ {
			tile := grid[gridRow][gridCol]

			// Copy inner part (without borders)
			for tileRow := 1; tileRow < tileSize-1; tileRow++ {
				for tileCol := 1; tileCol < tileSize-1; tileCol++ {
					resultRow := gridRow*innerSize + (tileRow - 1)
					resultCol := gridCol*innerSize + (tileCol - 1)
					result[resultRow][resultCol] = tile.data[tileRow][tileCol]
				}
			}
		}
	}

	return result
}

// Sea monster pattern
var seaMonster = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

// Check if sea monster exists at position
func checkSeaMonster(image [][]rune, row, col int) bool {
	for mr, line := range seaMonster {
		for mc, char := range line {
			if char == '#' {
				if row+mr >= len(image) || col+mc >= len(image[0]) {
					return false
				}
				if image[row+mr][col+mc] != '#' {
					return false
				}
			}
		}
	}
	return true
}

// Mark sea monster at position
func markSeaMonster(image [][]rune, row, col int) {
	for mr, line := range seaMonster {
		for mc, char := range line {
			if char == '#' {
				image[row+mr][col+mc] = 'O'
			}
		}
	}
}

// Find and mark all sea monsters
func findSeaMonsters(image [][]rune) int {
	count := 0
	monsterHeight := len(seaMonster)
	monsterWidth := len(seaMonster[0])

	for row := 0; row <= len(image)-monsterHeight; row++ {
		for col := 0; col <= len(image[0])-monsterWidth; col++ {
			if checkSeaMonster(image, row, col) {
				markSeaMonster(image, row, col)
				count++
			}
		}
	}

	return count
}

// Count remaining # characters
func countRoughWaters(image [][]rune) int {
	count := 0
	for _, row := range image {
		for _, char := range row {
			if char == '#' {
				count++
			}
		}
	}
	return count
}

// Create Image from rune array
func makeImageFromRunes(data [][]rune) *Image {
	return &Image{title: "assembled", data: data}
}

// Try all orientations to find sea monsters
func (p *Picture) FindRoughWaters() int {
	assembled := p.AssembleImage()
	img := makeImageFromRunes(assembled)

	// Try all 8 orientations
	for _, orientation := range img.GetOrientations() {
		// Make a copy to mark monsters
		imageCopy := make([][]rune, len(orientation.data))
		for i := range orientation.data {
			imageCopy[i] = make([]rune, len(orientation.data[i]))
			copy(imageCopy[i], orientation.data[i])
		}

		monstersFound := findSeaMonsters(imageCopy)
		if monstersFound > 0 {
			return countRoughWaters(imageCopy)
		}
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	imageData := []string{}
	for scanner.Scan() {
		imageData = append(imageData, scanner.Text())
	}

	picture := MakePicture(imageData)

	// Part 1
	corners := picture.FindCorners()
	total := 1
	for _, corner := range corners {
		var num int
		fmt.Sscanf(corner.title, "Tile %d:", &num)
		total *= num
	}
	fmt.Println("Part 1:", total)

	// Part 2
	roughWaters := picture.FindRoughWaters()
	fmt.Println("Part 2:", roughWaters)
}
