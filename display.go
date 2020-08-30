package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	ColorBlack = "\033[30m"
	ColorWhite = "\033[37m"

	ColorRed     = "\033[31m" // 625 - 740
	ColorYellow  = "\033[33m" // 565 - 590
	ColorGreen   = "\033[32m" // 520 - 565
	ColorCyan    = "\033[36m" // 500 - 520
	ColorBlue    = "\033[34m" // 435 - 500
	ColorMagenta = "\033[35m" // 380 - 435

	ColorReset = "\033[0m"

	Clear = "\033[2J"

	Corner         = "+"
	HorizontalLine = "-"
	VerticalLine   = "|"
)

func RenderInit(e Engine) {
	ClearScreen()

	// Jump(2, e.Height/2-3)
	// fmt.Print("==")

	Jump(e.Width+6, 7)
	fmt.Printf("Count: %d", e.Hits)
	Jump(e.Width+6, 9)
	fmt.Printf("Impact Parameter (b): %5.2f", 0.0)
	Jump(e.Width+6, 10)
	fmt.Printf("Scattering Angle (θ): %4.2f", 0.0)
	Jump(e.Width+6, 11)
	fmt.Printf("Detector: %1d", 0)

	Jump(e.Width+6, 2)
	fmt.Print("Scattering Angle's for two")
	Jump(e.Width+6, 3)
	fmt.Print("solid spheres elastic collisions")
	Jump(e.Width+6, 4)
	fmt.Print("at random impact parameters.")

	RenderBorder(e)
	RenderDetector(e)
}

func Jump(x, y int) {
	os.Stdout.WriteString(fmt.Sprintf("\033[%d;%df", y, x))
}

func ClearScreen() {
	fmt.Println(Clear)
}

func Render(e Engine) {
	RenderDetector(e)
	RenderBalls(e)
	Jump(e.Width, e.Height/2+3)
	fmt.Println()
}

func UnrenderBalls(e Engine) {
	for _, c := range e.Balls {
		Jump(int(math.Round(c.P.X)+2), int(math.Round(c.P.Y/2))+2)
		fmt.Println(" ")
	}
}

func RenderBalls(e Engine) {
	for _, c := range e.Balls {
		Jump(int(math.Round(c.P.X)+2), int(math.Round(c.P.Y/2))+2)
		fmt.Println("*")
	}
}

func RenderBorder(e Engine) {
	Jump(1, 1)
	fmt.Printf("+%s+", strings.Repeat("-", int(e.Width+1)))
	for y := 1; y < (e.Height/2)+2; y++ {
		Jump(1, 1+y)
		fmt.Print("|")
		Jump(e.Width+3, 1+y)
		fmt.Print("|")
	}
	Jump(1, (e.Height/2)+3)
	fmt.Printf("+%s+", strings.Repeat("-", int(e.Width+1)))
}

func DetectorColor(i int, e Engine) string {
	if e.LastDetectorHit == i {
		return ColorRed
	} else {
		return ColorReset
	}
}

func RenderDetector(e Engine) {
	startX := 35
	startY := 3
	Jump(startX+5, startY)
	fmt.Print(DetectorColor(2, e), "_____")
	Jump(startX+3, startY+1)
	fmt.Print(DetectorColor(3, e), ".'", DetectorColor(1, e), "     `.")
	Jump(startX+1, startY+2)
	fmt.Print(DetectorColor(3, e), ".'", DetectorColor(1, e), "         `.")
	Jump(startX, startY+3)
	fmt.Print(DetectorColor(3, e), "/", DetectorColor(1, e), "             \\")
	Jump(startX, startY+4)
	fmt.Print(DetectorColor(0, e), "               |")
	Jump(startX, startY+5)
	fmt.Print(DetectorColor(0, e), "               |")
	Jump(startX, startY+6)
	fmt.Print(DetectorColor(0, e), "               |")
	Jump(startX, startY+7)
	fmt.Print(DetectorColor(5, e), "\\", DetectorColor(7, e), "             /")
	Jump(startX+1, startY+8)
	fmt.Print(DetectorColor(5, e), "`.", DetectorColor(7, e), "         .'")
	Jump(startX+3, startY+9)
	fmt.Print(DetectorColor(5, e), "`.", DetectorColor(6, e), "_____", DetectorColor(7, e), ".'")
	fmt.Print(ColorReset)

	Jump(startX+17, startY+5)
	fmt.Print(e.Detector[0])
	Jump(startX+14, startY+1)
	fmt.Print(e.Detector[1])
	Jump(startX+7, startY-1)
	fmt.Print(e.Detector[2])
	Jump(startX, startY+1)
	fmt.Print(e.Detector[3])

	Jump(startX, startY+9)
	fmt.Print(e.Detector[5])
	Jump(startX+7, startY+10)
	fmt.Print(e.Detector[6])
	Jump(startX+14, startY+9)
	fmt.Print(e.Detector[7])

}

// |                                        _____          |
// |                                      .'     `.         |
// |                                    .'         `.        |
// |                                   /             \       |
// |                                                  |  10    |
// |                                          *       |  10    |
// |                                                  |     |
// |                                   \             /         |
// |                                    `.         .'         |
// |                                      `._____.'          |
// |                                    8           8        |
