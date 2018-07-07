# Disclaimer
In this project I wanted to showcase some image drawing and cli tool building.<br><br>
If you build this programm you have a binary called "xy". You can place that in your path and run:
```
$ xy checkpoint --help
```

# Installation
- go get github.com/zkynet/xy

# Test 
no tests dues to time limit .. maybe I'll add some later.

# Help
```
$ go run main.go checkpoint  --help
```

# Usage without building
 ```
 $ go run main.go checkpoint --rect-bottom-left-x 0 --rect-bottom-left-y 0 --rect-height 100 --rect-width 100 --point-x 50 --point-y 50 --img-out 1.png --bool-out
 ```

# Usage after building
```
$ xy checkpoint --rect-bottom-left-x 0 --rect-bottom-left-y 0 --rect-height 100 --rect-width 100 --point-x 50 --point-y 50 --img-out 1.png --bool-out
```

# References
- https://en.wikipedia.org/wiki/Cartesian_coordinate_system
- https://martin-thoma.com/how-to-check-if-a-point-is-inside-a-rectangle/

# The core of it
```

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

	return rectArea == (ABP + BCP + CDP + DAP)

}

```



# Results
- a message or boolean indicating if the point is inside or outside of the rectangle
- optional image showing the rectangle and where the point is

# Rendered image
![alt text](draw1.png) ![alt text](draw.png)  ![alt text](draw2.png)

