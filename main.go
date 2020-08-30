package main

import (
	"math/rand"
	"time"
)

func main() {
	// Scattering()
	RandomMotion()
}

func Scattering() {

	rand.Seed(time.Now().Unix())
	e := NewEngine(50, 6, 0)

	v := 50.0
	e.Balls = append(e.Balls, &Circle{P: Vec{X: 40, Y: 3}, V: Vec{X: 0, Y: 0}, Mass: 1000000, Name: "target", Radius: .5})
	e.Balls = append(e.Balls, &Circle{P: Vec{X: 1, Y: 2}, V: Vec{X: v, Y: 0}, Mass: 1, Radius: .5})

	Render(e)
	RenderInit(e)

	fps := float64(1000)
	frame := time.Second / time.Duration(fps)
	step := 0
	for {
		if step%1000 == 0 {
			UnrenderBalls(e)
			e.Balls[1].P = Vec{X: 1, Y: 2 + float64(step)/10000.0}
			e.Balls[1].V = Vec{X: v, Y: 0}
		}

		start := time.Now()
		UnrenderBalls(e)
		e.Step(1.0 / fps)
		e.HandleCollisions()
		Render(e)

		elapsed := time.Now().Sub(start)
		time.Sleep(frame - elapsed)
		step++
	}

}

func RandomMotion() {
	rand.Seed(time.Now().Unix())

	e := NewEngine(80, 25, 5)
	// e.Balls = append(e.Balls, &Circle{P: Vec{X: 1, Y: 10}, V: Vec{X: 5, Y: 0}})
	// e.Balls = append(e.Balls, &Circle{P: Vec{X: 10, Y: 9.5}, V: Vec{X: -5, Y: 0}})

	Render(e)
	RenderInit(e)

	fps := float64(45)
	frame := time.Second / time.Duration(fps)
	for {
		start := time.Now()
		UnrenderBalls(e)
		e.Step(1.0 / fps)
		e.HandleCollisions()
		Render(e)

		elapsed := time.Now().Sub(start)
		time.Sleep(frame - elapsed)
	}
}
