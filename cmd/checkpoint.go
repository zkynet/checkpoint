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
	BLX float64
	BLY float64
	TLX float64
	TLY float64

	// right side
	BRX float64
	BRY float64
	TRX float64
	TRY float64
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
	rect := &Rect{
		BLX: bottomLeftx,
		BLY: bottomLefty,
	}
	rect.TLX = bottomLeftx
	rect.TLY = bottomLefty + height

	rect.TRX = bottomLeftx + width
	rect.TRY = bottomLefty + height

	rect.BRX = bottomLeftx + width
	rect.BRY = bottomLefty

	fmt.Println("Rectangle")
	fmt.Println(rect)
	fmt.Println("Top Left X:", rect.TLX)
	fmt.Println("Top Left Y:", rect.TLY)
	fmt.Println("Bottom Left X:", rect.BLX)
	fmt.Println("Bottom Left Y:", rect.BLY)

	fmt.Println("Top Right X:", rect.TRX)
	fmt.Println("Top Right Y:", rect.TRY)
	fmt.Println("Bottom Right X:", rect.BRX)
	fmt.Println("Bottom Right Y:", rect.BRY)

	img = image.NewRGBA(image.Rect(0, 0, 50, 50))

	//col = color.RGBA{255, 0, 0, 255} // Red
	//RectDraw(15, 15, 20, 20)
	col = color.RGBA{0, 255, 0, 255} // Green
	RectDraw(round(rect.BLX), round(rect.BLY), round(rect.TRX), round(rect.TRY))
	col = color.RGBA{255, 0, 0, 255} // Red
	RectDraw(round(x), round(y), round(x), round(y))

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

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
	// is called directly, e.g.:
	// checkpointCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
