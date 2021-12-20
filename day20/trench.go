package day20

import (
	"fmt"
	"strings"
)

type PixelMap string
func (pm *PixelMap) pixel(idx int) uint8 {
	return (*pm)[idx]
}

type Image struct {
	horizonPixel uint8
	pixels [][]uint8
}

func (i Image) definedHeight() int {
	return len(i.pixels)
}

func (i Image) definedWidth() int {
	return len(i.pixels[0])
}


func (i Image) String() string {
	var lines []string
	for _, row := range i.pixels {
		line := ""
		for _, pixel := range row {
			line += string(pixel)
		}
		lines = append(lines, line)
	}

	lines = append(lines, fmt.Sprintf("(invisible pixels are %s)", string(i.horizonPixel)))

	return strings.Join(lines, "\n")
}

func (i Image) pixelAt(row, col int) uint8 {
	if row < 0 || col < 0 || row >= i.definedHeight() || col >= i.definedWidth() {
		return i.horizonPixel
	} else {
		return i.pixels[row][col]
	}
}

func NewImageFromLines(lines []string) *Image {
	newImage := Image{
		horizonPixel: '.',
		pixels: make([][]uint8, len(lines)),
	}
	for i := 0; i < len(lines); i++ {
		newImage.pixels[i] = make([]uint8, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			newImage.pixels[i][j] = lines[i][j]
		}
	}
	return &newImage
}

func square() [][]int {
	return [][]int{
		{-1,-1}, {-1,0}, {-1,1},
		{ 0,-1}, { 0,0}, { 0,1},
		{ 1,-1}, { 1,0}, { 1,1},
	}
}

func (i *Image) enhance(pixelMap PixelMap) *Image {
	newHeight := i.definedHeight() + 2
	newWidth := i.definedWidth() + 2
	newPixels := make([][]uint8, newHeight)
	for newRow := 0; newRow < newWidth; newRow++ {
		newPixels[newRow] = make([]uint8, newWidth)
		for newCol := 0; newCol < newHeight; newCol++ {
			mask := square()
			value := 0
			power := 1
			for bit := 8; bit >= 0; bit-- {
				sourceRow := newRow + mask[bit][0] - 1
				sourceCol := newCol + mask[bit][1] - 1
				if i.pixelAt(sourceRow, sourceCol) == '#' {
					value += power
				}
				power *= 2
			}
			newPixels[newRow][newCol] = pixelMap.pixel(value)
		}
	}

	var newHorizonPixel uint8
	if i.horizonPixel == '.' {
		newHorizonPixel = pixelMap.pixel(0)
	} else {
		newHorizonPixel = pixelMap.pixel(511)
	}

	return &Image{
		horizonPixel: newHorizonPixel,
		pixels: newPixels,
	}
}

func (i Image) pixelsLit() (lit int, infinite bool) {
	if i.horizonPixel == '#' {
		return 0, true
	}

	for row := 0; row < i.definedHeight(); row++ {
		for col := 0; col < i.definedHeight(); col++ {
			if i.pixelAt(row, col) == '#' {
				lit += 1
			}
		}
	}
	return lit, false
}

func enhanceMany(pm PixelMap, source *Image, times int) *Image {
	current := source
	for i := 0; i < times; i++ {
		current = current.enhance(pm)
	}
	return current
}

func describeLitPixels(rounds int, image *Image) string {
	lit, infinity := image.pixelsLit()
	if infinity {
		panic(fmt.Sprintf("Infinite amount of pixels lit after round %d!", rounds))
	} else {
		return fmt.Sprintf("Pixels lit after %d enhancements: %d", rounds, lit)
	}
}

func solve(lines []string, rounds int) string {
	pm := PixelMap(lines[0])
	image := NewImageFromLines(lines[2:])

	enhanced := enhanceMany(pm, image, rounds)
	return describeLitPixels(rounds, enhanced)
}

func Part1(lines []string) string {
	return solve(lines, 2)
}

func Part2(lines []string) string {
	return solve(lines, 50)
}
