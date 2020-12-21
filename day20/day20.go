package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Image struct {
	title string
	data  [][]rune
}

type Picture struct {
	tiles []*Image
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
	corners := picture.FindCorners()
	total := 1
	for _, corner := range corners {
		var num int
		fmt.Sscanf(corner.title, "Tile %d:", &num)
		total *= num
	}
	fmt.Println("Part 1:", total)
}
