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

	RenderBorder(e)
}

func Jump(x, y int) {
	os.Stdout.WriteString(fmt.Sprintf("\033[%d;%df", y, x))
}

func ClearScreen() {
	fmt.Println(Clear)
}

func Render(e Engine) {
	RenderBalls(e)
	Jump(e.Width, e.Height+3)
	fmt.Println()
}

func UnrenderBalls(e Engine) {
	for _, c := range e.Balls {
		Jump(int(math.Round(c.P.X)+2), int(math.Round(c.P.Y))+2)
		fmt.Println(" ")
	}
}

func RenderBalls(e Engine) {
	for _, c := range e.Balls {
		Jump(int(math.Round(c.P.X)+2), int(math.Round(c.P.Y))+2)
		fmt.Println("*")
	}
}

func RenderBorder(e Engine) {
	Jump(1, 1)
	fmt.Printf("+%s+", strings.Repeat("-", int(e.Width+1)))
	for y := 1; y < e.Height+2; y++ {
		Jump(1, 1+y)
		fmt.Print("|")
		Jump(e.Width+3, 1+y)
		fmt.Print("|")
	}
	Jump(1, e.Height+3)
	fmt.Printf("+%s+", strings.Repeat("-", int(e.Width+1)))
}
