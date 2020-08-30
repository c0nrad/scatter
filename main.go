package main

import (
	"math/rand"
	"time"
)

func main() {
	Scattering()
}

func Scattering() {
	rand.Seed(time.Now().Unix())
	e := NewEngine(52, 22)

	e.Balls = append(e.Balls, &Circle{P: Vec{X: 40, Y: 12}, V: Vec{X: 0, Y: 0}, Mass: 1000000, Name: "target", Radius: .5})

	Render(e)
	RenderInit(e)

	fps := float64(2000)
	frame := time.Second / time.Duration(fps)
	step := 0
	for {
		e.Balls[0].V.X = 0
		e.Balls[0].V.Y = 0
		e.Balls[0].P.X = 40
		e.Balls[0].P.Y = 12
		if step%200 == 0 {
			e.Fire(rand.Float64()*1.5 - .75)
		}

		start := time.Now()
		UnrenderBalls(e)
		e.Step(1.0 / fps)
		e.HandleCollisions()
		e.HandleDetectors()
		Render(e)

		elapsed := time.Now().Sub(start)
		time.Sleep(frame - elapsed)
		step++

		if step == 40000 {
			Plot(e.BThetaMap)
		}
	}
}
