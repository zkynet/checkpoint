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

// checkpointCmd represents the checkrect command
var checkpointCmd = &cobra.Command{
	Use:   "checkpoint",
	Short: "Check if a point falls within a 2D boundry of the given rectangle",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		checkPoint()
	},
}

func checkPoint() {
	rect := Rect{
		AX: bottomLeftx,
		AY: bottomLefty,
	}
	rect.BX = bottomLeftx
	rect.BY = bottomLefty + height
	rect.CX = bottomLeftx + width
	rect.CY = bottomLefty + height
	rect.DX = bottomLeftx + width
	rect.DY = bottomLefty

	point := Point{
		X: x,
		Y: y,
	}
	fmt.Println("Point:")
	fmt.Println(point)
	fmt.Println("Rectangle:")
	fmt.Println(rect)
	fmt.Println("height:", height)
	fmt.Println("width:", width)

	fmt.Println("Top Left X:", rect.BX)
	fmt.Println("Top Left Y:", rect.BY)
	fmt.Println("Bottom Left X:", rect.AX)
	fmt.Println("Bottom Left Y:", rect.AY)

	fmt.Println("Top Right X:", rect.CX)
	fmt.Println("Top Right Y:", rect.CY)
	fmt.Println("Bottom Right X:", rect.DX)
	fmt.Println("Bottom Right Y:", rect.DY)

	// calculate area
	// (bottom left y - top right y ) * (bottom right x - top left x) + (top left y - bottom right y) * (bottom left x - top right x)
	rectArea := 0.5 * math.Abs(((rect.AY-rect.CY)*(rect.DX-rect.BX))+((rect.BY-rect.DY)*(rect.AX-rect.CX)))
	fmt.Println(rectArea)

	ABP := 0.5 * math.Abs((rect.AX*(rect.BY-point.Y) + rect.BX*(point.Y-rect.AY) + point.X*(rect.AY-rect.BY)))
	fmt.Println(ABP)
	BCP := 0.5 * math.Abs((rect.BX*(rect.CY-point.Y) + rect.CX*(point.Y-rect.BY) + point.X*(rect.BY-rect.CY)))
	fmt.Println(BCP)
	CDP := 0.5 * math.Abs((rect.CX*(rect.DY-point.Y) + rect.DX*(point.Y-rect.CY) + point.X*(rect.CY-rect.DY)))
	fmt.Println(CDP)
	DAP := 0.5 * math.Abs((rect.DX*(rect.AY-point.Y) + rect.AX*(point.Y-rect.DY) + point.X*(rect.DY-rect.AY)))
	fmt.Println(DAP)

	img = image.NewRGBA(image.Rect(0, 0, round(rect.AX)+round(width)+10, round(rect.AY)+round(height)+10))
	col = color.RGBA{0, 255, 0, 255} // Green
	RectDraw(round(rect.AX), round(rect.AY), round(rect.CX), round(rect.CY))
	col = color.RGBA{255, 0, 0, 255} // Red
	HLine(round(x), round(y), round(x)+5)
	HLine(round(x)-5, round(y), round(x))
	VLine(round(x), round(y), round(y)+5)
	VLine(round(x), round(y)-5, round(y))

	col = color.RGBA{255, 0, 255, 255} // Red
	VLine(round(rect.AX), round(rect.AY), round(rect.AY))

	imgR := imaging.FlipV(img)
	f, err := os.Create("draw.png")
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

	checkpointCmd.Flags().Float64VarP(&bottomLeftx, "rect-bottom-left-x", "", 1, "Bottom Left point of the rectangle")
	checkpointCmd.Flags().Float64VarP(&bottomLefty, "rect-bottom-left-y", "", 1, "Bottom Left point of the rectangle")

	checkpointCmd.Flags().Float64VarP(&height, "rect-height", "H", 1, "Height of the tectangle")
	checkpointCmd.Flags().Float64VarP(&width, "rect-width", "W", 1, "Width of the tectangle")

	checkpointCmd.Flags().Float64VarP(&x, "point-X", "X", 1, "Point X")
	checkpointCmd.Flags().Float64VarP(&y, "point-Y", "Y", 1, "Point Y")

	rootCmd.AddCommand(checkpointCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkpointCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called direcBY, e.g.:
	// checkpointCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
