package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Circle struct {
	P, V      Vec
	PreviousP Vec

	B      float64
	Mass   float64
	Steps  int
	Radius float64

	Name string

	InDetector bool
}

type Engine struct {
	Balls []*Circle

	Time          float64
	Width, Height int

	DetectorRadius  int
	LastDetectorHit int
	Detector        map[int]int
	Hits            int
	BThetaMap       map[float64]float64
}

func NewEngine(width, height int) Engine {
	e := Engine{Width: width, Height: height, Time: 0, DetectorRadius: 8,
		BThetaMap:       make(map[float64]float64),
		Detector:        make(map[int]int),
		LastDetectorHit: 9,
	}

	return e
}

func (e *Engine) Fire(b float64) {
	c := Circle{}
	c.P.X = 2
	c.P.Y = e.Balls[0].P.Y - b
	c.V.X = 99 + 2*rand.Float64()
	c.Mass = 1
	c.B = b
	e.Balls = append(e.Balls, &c)
}

func (e *Engine) Step(dt float64) {
	for _, c := range e.Balls {
		c.PreviousP = c.P
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

func (e *Engine) HandleDetectors() {
	if len(e.Balls) < 2 {
		return
	}
	target := e.Balls[0]
	out := []*Circle{target}

	for i := 1; i < len(e.Balls); i++ {
		c := e.Balls[i]
		if !c.InDetector {
			out = append(out, c)
			if c.P.DistanceSquared(target.P) < float64(e.DetectorRadius*e.DetectorRadius) {
				c.InDetector = true
				continue
			}
			continue
		}

		if c.PreviousP.DistanceSquared(target.P) < float64(e.DetectorRadius*e.DetectorRadius) &&
			!(c.P.DistanceSquared(target.P) < float64(e.DetectorRadius*e.DetectorRadius)) {

			diff := c.P.Minus(target.P)
			theta := math.Atan2(-diff.Y, diff.X)
			if theta < 0 {
				theta += 2 * math.Pi
			}

			if theta < math.Pi/8 || theta >= math.Pi*15/8 {
				e.LastDetectorHit = 0
			} else if theta < math.Pi*3/8 && theta >= math.Pi*1/8 {
				e.LastDetectorHit = 1
			} else if theta < math.Pi*5/8 && theta >= math.Pi*3/8 {
				e.LastDetectorHit = 2
			} else if theta < math.Pi*7/8 && theta >= math.Pi*5/8 {
				e.LastDetectorHit = 3
			} else if theta < math.Pi*9/8 && theta >= math.Pi*7/8 {
				e.LastDetectorHit = 4
			} else if theta < math.Pi*11/8 && theta >= math.Pi*9/8 {
				e.LastDetectorHit = 5
			} else if theta < math.Pi*13/8 && theta >= math.Pi*11/8 {
				e.LastDetectorHit = 6
			} else if theta < math.Pi*15/8 && theta >= math.Pi*13/8 {
				e.LastDetectorHit = 7
			}
			e.BThetaMap[c.B] = theta
			e.Detector[e.LastDetectorHit]++
			e.Hits++

			Jump(e.Width+6, 7)
			fmt.Printf("Count: %d", e.Hits)

			Jump(e.Width+6, 9)
			fmt.Printf("Impact Parameter (b): %5.2f", c.B)
			Jump(e.Width+6, 10)
			fmt.Printf("Scattering Angle (θ): %4.2f", theta)
			Jump(e.Width+6, 11)
			fmt.Printf("Detector: %1d", e.LastDetectorHit)

			// leaving detector
			c.P.X = 1
			c.P.Y = 1
			c.V.X = 0
			c.V.Y = 0
		} else {
			out = append(out, c)
		}
	}

	e.Balls = out
}
