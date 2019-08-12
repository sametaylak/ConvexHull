package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"

	"github.com/sametaylak/convex-hull/vector"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowTitle  = "Convex Hull"
	WindowWidth  = 800
	WindowHeight = 600
	FrameRate    = 60
	NumPoints    = 100

	MaxX = WindowWidth - 100
	MinX = 100
	MaxY = WindowHeight - 100
	MinY = 100
)

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	sdl.Do(func() {
		window, err = sdl.CreateWindow(
			WindowTitle,
			sdl.WINDOWPOS_UNDEFINED,
			sdl.WINDOWPOS_UNDEFINED,
			WindowWidth,
			WindowHeight,
			sdl.WINDOW_OPENGL,
		)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer func() {
		sdl.Do(func() {
			window.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer, err = sdl.CreateRenderer(
			window,
			-1,
			sdl.RENDERER_ACCELERATED,
		)
	})
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		sdl.Do(func() {
			renderer.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer.Clear()
	})

	hull := []sdl.Point{}
	points := make([]vector.Vector, NumPoints)
	for i := range points {
		points[i] = vector.Vector{
			X: int32(rand.Intn(MaxX-MinX) + MinX),
			Y: int32(rand.Intn(MaxY-MinY) + MinY),
			Z: 0,
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})

	currentPoint := points[0]
	hull = append(hull, sdl.Point{
		X: currentPoint.X,
		Y: currentPoint.Y,
	})
	nextPoint := points[1]
	checkPointIdx := 2
	checkPoint := points[2]

	running := true
	for running {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					running = false
				}
			}

			renderer.Clear()
			renderer.SetDrawColor(0x00, 0x00, 0x00, 0x20)
			renderer.FillRect(&sdl.Rect{0, 0, WindowWidth, WindowHeight})

			for i := range points {
				renderer.SetDrawColor(0xff, 0xff, 0xff, 0xff)
				renderer.DrawPoint(points[i].X, points[i].Y)
			}

			if len(hull) > 0 {
				renderer.SetDrawColor(0x00, 0x00, 0xff, 0xff)
				renderer.DrawLines(hull)
			}

			renderer.SetDrawColor(0x00, 0xff, 0x00, 0xff)
			renderer.DrawLine(
				currentPoint.X,
				currentPoint.Y,
				nextPoint.X,
				nextPoint.Y,
			)

			renderer.SetDrawColor(0xff, 0x00, 0x00, 0xff)
			renderer.DrawLine(
				currentPoint.X,
				currentPoint.Y,
				checkPoint.X,
				checkPoint.Y,
			)

			renderer.Present()

			vectorA := vector.Sub(currentPoint, nextPoint)
			vectorB := vector.Sub(checkPoint, currentPoint)
			crossProduct := vector.CrossProduct(vectorA, vectorB)
			if crossProduct.Z < 0 {
				nextPoint = checkPoint
			}

			checkPointIdx += 1
			if checkPointIdx == len(points) {
				if nextPoint == points[0] {
					fmt.Println("Done")
					running = false
				}

				hull = append(hull, sdl.Point{
					X: nextPoint.X,
					Y: nextPoint.Y,
				})

				currentPoint = nextPoint
				nextPoint = points[0]
				checkPointIdx = 0
			}
			checkPoint = points[checkPointIdx]

			sdl.Delay(1000 / FrameRate)
		})
	}

	return 0
}

func main() {
	var exitcode int
	sdl.Main(func() {
		exitcode = run()
	})
	os.Exit(exitcode)
}
