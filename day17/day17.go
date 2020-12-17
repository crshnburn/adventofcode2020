package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Grid struct {
	layers     [][][]bool
	layerSize  int
	layerDepth int
}

type Cube struct {
	cubes      [][][][]bool
	layerSize  int
	layerDepth int
	layerCount int
}

func (g *Grid) Neighbours(x, y, z int) int {
	count := 0
	for z1 := -1; z1 <= 1; z1++ {
		for y1 := -1; y1 <= 1; y1++ {
			for x1 := -1; x1 <= 1; x1++ {
				xloc := x + x1
				yloc := y + y1
				zloc := z + z1
				if xloc < 0 || yloc < 0 || zloc < 0 || (xloc == x && yloc == y && zloc == z) || zloc >= g.layerDepth || yloc >= g.layerSize || xloc >= g.layerSize {
					//do nothing as either initialised to false or our location
				} else {
					if g.layers[zloc][yloc][xloc] {
						count++
					}
				}
			}
		}
	}
	return count
}

func (g *Cube) Neighbours(x, y, z, w int) int {
	count := 0
	for w1 := -1; w1 <= 1; w1++ {
		for z1 := -1; z1 <= 1; z1++ {
			for y1 := -1; y1 <= 1; y1++ {
				for x1 := -1; x1 <= 1; x1++ {
					xloc := x + x1
					yloc := y + y1
					zloc := z + z1
					wloc := w + w1
					if xloc < 0 || yloc < 0 || zloc < 0 || wloc < 0 || (xloc == x && yloc == y && zloc == z && wloc == w) || zloc >= g.layerDepth || yloc >= g.layerSize || xloc >= g.layerSize || wloc >= g.layerCount {
						//do nothing as either initialised to false or our location
					} else {
						if g.cubes[wloc][zloc][yloc][xloc] {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func (g *Grid) String() {
	for _, layer := range g.layers {
		for _, row := range layer {
			for _, val := range row {
				if val {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func (g *Grid) Scale() *Grid {
	newSize := g.layerSize + 2
	newDepth := g.layerDepth + 2
	newLayers := MakeNewView(newSize, newDepth)
	for z := 1; z < newDepth-1; z++ {
		for y := 1; y < newSize-1; y++ {
			for x := 1; x < newSize-1; x++ {
				newLayers[z][y][x] = g.layers[z-1][y-1][x-1]
			}
		}
	}
	return &Grid{layers: newLayers, layerSize: newSize, layerDepth: newDepth}
}

func (g *Cube) Scale() *Cube {
	newSize := g.layerSize + 2
	newDepth := g.layerDepth + 2
	newCount := g.layerCount + 2
	newLayers := MakeNewCube(newSize, newDepth, newCount)
	for w := 1; w < newCount-1; w++ {
		for z := 1; z < newDepth-1; z++ {
			for y := 1; y < newSize-1; y++ {
				for x := 1; x < newSize-1; x++ {
					newLayers[w][z][y][x] = g.cubes[w-1][z-1][y-1][x-1]
				}
			}
		}
	}
	return &Cube{cubes: newLayers, layerSize: newSize, layerDepth: newDepth, layerCount: newCount}
}

func (g *Grid) Cycle() *Grid {
	updatedLayers := MakeNewView(g.layerSize, g.layerDepth)
	for z := 0; z < g.layerDepth; z++ {
		for y := 0; y < g.layerSize; y++ {
			for x := 0; x < g.layerSize; x++ {
				count := g.Neighbours(x, y, z)
				if g.layers[z][y][x] {
					if count == 2 || count == 3 {
						updatedLayers[z][y][x] = true
					} else {
						updatedLayers[z][y][x] = false
					}
				} else {
					if count == 3 {
						updatedLayers[z][y][x] = true
					} else {
						updatedLayers[z][y][x] = false
					}
				}
			}
		}
	}
	return &Grid{layers: updatedLayers, layerSize: g.layerSize, layerDepth: g.layerDepth}
}

func (g *Cube) Cycle() *Cube {
	updatedLayers := MakeNewCube(g.layerSize, g.layerDepth, g.layerCount)
	for w := 0; w < g.layerCount; w++ {
		for z := 0; z < g.layerDepth; z++ {
			for y := 0; y < g.layerSize; y++ {
				for x := 0; x < g.layerSize; x++ {
					count := g.Neighbours(x, y, z, w)
					if g.cubes[w][z][y][x] {
						if count == 2 || count == 3 {
							updatedLayers[w][z][y][x] = true
						} else {
							updatedLayers[w][z][y][x] = false
						}
					} else {
						if count == 3 {
							updatedLayers[w][z][y][x] = true
						} else {
							updatedLayers[w][z][y][x] = false
						}
					}
				}
			}
		}
	}
	return &Cube{cubes: updatedLayers, layerSize: g.layerSize, layerDepth: g.layerDepth, layerCount: g.layerCount}
}

func (g *Grid) Count() int {
	count := 0
	for z := 0; z < g.layerDepth; z++ {
		for y := 0; y < g.layerSize; y++ {
			for x := 0; x < g.layerSize; x++ {
				if g.layers[z][y][x] {
					count++
				}
			}
		}
	}
	return count
}

func (g *Cube) Count() int {
	count := 0
	for w := 0; w < g.layerCount; w++ {
		for z := 0; z < g.layerDepth; z++ {
			for y := 0; y < g.layerSize; y++ {
				for x := 0; x < g.layerSize; x++ {
					if g.cubes[w][z][y][x] {
						count++
					}
				}
			}
		}
	}
	return count
}

func MakeNewView(size, depth int) [][][]bool {
	var newLayers [][][]bool
	for z := 0; z < depth; z++ {
		var layer [][]bool
		for y := 0; y < size; y++ {
			layer = append(layer, make([]bool, size))
		}
		newLayers = append(newLayers, layer)
	}
	return newLayers
}

func MakeNewCube(size, depth, count int) [][][][]bool {
	var newView [][][][]bool
	for w := 0; w < count; w++ {
		var cube [][][]bool
		for z := 0; z < depth; z++ {
			var layer [][]bool
			for y := 0; y < size; y++ {
				layer = append(layer, make([]bool, size))
			}
			cube = append(cube, layer)
		}
		newView = append(newView, cube)
	}
	return newView
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var layer [][]bool
	for scanner.Scan() {
		var row []bool
		for _, char := range scanner.Text() {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		layer = append(layer, row)
	}

	layers := [][][]bool{layer}

	grid := &Grid{layers: layers, layerSize: len(layer), layerDepth: 1}
	// grid.String()
	for i := 0; i < 6; i++ {
		grid = grid.Scale()
		grid = grid.Cycle()
	}
	fmt.Println("Part 1:", grid.Count())

	cube := &Cube{cubes: [][][][]bool{layers}, layerSize: len(layer), layerDepth: 1, layerCount: 1}
	for i := 0; i < 6; i++ {
		cube = cube.Scale()
		cube = cube.Cycle()
	}
	fmt.Println("Part 2:", cube.Count())
}
