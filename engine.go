package main

import (
	"math/rand"
)

type Circle struct {
	P, V   Vec
	Mass   float64
	Steps  int
	Radius float64

	Name string
}

type Engine struct {
	Balls []*Circle

	Time float64

	Width, Height int
}

func NewEngine(width, height, count int) Engine {
	e := Engine{Width: width, Height: height, Time: 0}

	for i := 0; i < count; i++ {
		c := Circle{}
		c.P.X = float64(rand.Intn(width))
		c.P.Y = float64(rand.Intn(height))
		c.V.X = float64(rand.Intn(100))
		c.V.Y = float64(rand.Intn(100))
		c.Mass = float64(rand.Intn(10))
		e.Balls = append(e.Balls, &c)
	}
	return e
}

func (e *Engine) Step(dt float64) {
	for _, c := range e.Balls {
		c.P.X += dt * c.V.X
		c.P.Y += dt * c.V.Y

		if c.P.X > float64(e.Width) {
			c.P.X = 2*float64(e.Width) - c.P.X
			c.V.X *= -1
		}
		if c.P.Y > float64(e.Height) {
			c.P.Y = 2*float64(e.Height) - c.P.Y
			c.V.Y *= -1
		}
		if c.P.Y < 0 {
			c.P.Y = -c.P.Y
			c.V.Y *= -1
		}
		if c.P.X < 0 {
			c.P.X = -c.P.X
			c.V.X *= -1
		}

		c.Steps++
	}
	e.Time += dt
}

func (e *Engine) HandleCollisions() {
	for c1i := 0; c1i < len(e.Balls); c1i++ {
		for c2i := c1i + 1; c2i < len(e.Balls); c2i++ {
			c1 := e.Balls[c1i]
			c2 := e.Balls[c2i]

			if c1.P.DistanceSquared(c2.P) < (c1.Radius+c2.Radius)*(c1.Radius+c2.Radius) {
				// Sadly I don't understand these equations. But they work. Forgive me.
				// https://en.wikipedia.org/wiki/Elastic_collision
				v1 := c1.V.Minus(c1.P.Minus(c2.P).Scale(2 * c2.Mass / (c1.Mass + c2.Mass) * c1.V.Minus(c2.V).Dot(c1.P.Minus(c2.P)) / c1.P.Minus(c2.P).Dot(c1.P.Minus(c2.P))))
				v2 := c2.V.Minus(c2.P.Minus(c1.P).Scale(2 * c1.Mass / (c1.Mass + c2.Mass) * c2.V.Minus(c1.V).Dot(c2.P.Minus(c1.P)) / c1.P.Minus(c2.P).Dot(c1.P.Minus(c2.P))))
				c1.V = v1
				c2.V = v2
			}
		}
	}
}
