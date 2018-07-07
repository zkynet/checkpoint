// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
)

var img *image.RGBA
var col color.Color

// HLine draws a horizontal line
func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, col)
	}
}

// Rect draws a rectangle utilizing HLine() and VLine()
func RectDraw(x1, y1, x2, y2 int) {
	HLine(x1, y1, x2)
	HLine(x1, y2, x2)
	VLine(x1, y1, y2)
	VLine(x2, y1, y2)
}

type Rect struct {

	// left side
	AX float64
	AY float64
	BX float64
	BY float64

	// right side
	DX float64
	DY float64
	CX float64
	CY float64
}

type Point struct {
	X float64
	Y float64
}

var bottomLeftx float64
var bottomLefty float64

var height float64
var width float64

var x float64
var y float64

var imageOut string
var boolOnly bool

// checkpointCmd represents the checkrect command
var checkpointCmd = &cobra.Command{
	Use:   "checkpoint",
	Short: "Check if a point falls within a 2D boundry of the given rectangle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		if boolOnly {
			fmt.Println(checkPoint())
		} else {
			if checkPoint() {
				fmt.Println("The point is inside the rectangle boundries")
			} else {
				fmt.Println("The point is outside the rectangle boundries")
			}
		}

	},
}

func checkPoint() bool {
	rect := Rect{
		AX: bottomLeftx,
		AY: bottomLefty,
		BX: bottomLeftx,
		BY: bottomLefty + height,
		CX: bottomLeftx + width,
		CY: bottomLefty + height,
		DX: bottomLeftx + width,
		DY: bottomLefty,
	}

	point := Point{
		X: x,
		Y: y,
	}

	// Calculate the area of the original rectangle
	rectArea := 0.5 * math.Abs(((rect.AY-rect.CY)*(rect.DX-rect.BX))+((rect.BY-rect.DY)*(rect.AX-rect.CX)))
	// Calculate a rectangle area using our new point
	ABP := 0.5 * math.Abs((rect.AX*(rect.BY-point.Y) + rect.BX*(point.Y-rect.AY) + point.X*(rect.AY-rect.BY)))
	BCP := 0.5 * math.Abs((rect.BX*(rect.CY-point.Y) + rect.CX*(point.Y-rect.BY) + point.X*(rect.BY-rect.CY)))
	CDP := 0.5 * math.Abs((rect.CX*(rect.DY-point.Y) + rect.DX*(point.Y-rect.CY) + point.X*(rect.CY-rect.DY)))
	DAP := 0.5 * math.Abs((rect.DX*(rect.AY-point.Y) + rect.AX*(point.Y-rect.DY) + point.X*(rect.DY-rect.AY)))

	if imageOut != "" {
		makeImage(&rect)
	}

	return rectArea == (ABP + BCP + CDP + DAP)

}

func makeImage(rect *Rect) {
	// Draw the rectangle
	img = image.NewRGBA(image.Rect(0, 0, round(rect.AX)+round(width)+round(x)+100, round(rect.AY)+round(height)+round(y)+100))
	col = color.RGBA{0, 255, 0, 255} // Green
	RectDraw(round(rect.AX)+100, round(rect.AY)+100, round(rect.CX)+100, round(rect.CY)+100)

	// Draw the point
	col = color.RGBA{255, 0, 0, 255} // Red
	HLine(round(x)+100, round(y)+100, round(x)+105)
	HLine(round(x)+95, round(y)+100, round(x)+100)
	VLine(round(x)+100, round(y)+100, round(y)+105)
	VLine(round(x)+100, round(y)+95, round(y)+100)

	// flip the image for a correct x/y axis
	imgR := imaging.FlipV(img)
	f, err := os.Create(imageOut)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, imgR)
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func init() {
	checkpointCmd.Flags().Float64VarP(&bottomLeftx, "rect-bottom-left-x", "", 1, "Bottom Left X point of the rectangle")
	checkpointCmd.Flags().Float64VarP(&bottomLefty, "rect-bottom-left-y", "", 1, "Bottom Left Y point of the rectangle")
	checkpointCmd.Flags().Float64VarP(&height, "rect-height", "H", 1, "Height of the tectangle")
	checkpointCmd.Flags().Float64VarP(&width, "rect-width", "W", 1, "Width of the tectangle")
	checkpointCmd.Flags().Float64VarP(&x, "point-x", "X", 1, "Point X")
	checkpointCmd.Flags().Float64VarP(&y, "point-y", "Y", 1, "Point Y")
	checkpointCmd.Flags().StringVarP(&imageOut, "img-out", "i", "", "Output Image Path, if no image is specified the program will not make an image")
	checkpointCmd.Flags().BoolVarP(&boolOnly, "bool-out", "b", false, "Output bool instead of a message")

	rootCmd.AddCommand(checkpointCmd)
}
