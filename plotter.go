package main

import (
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func Plot(bThetaMap map[float64]float64) {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Scattering Impact Partameter vs Scattering Angle for Solid Spheres"
	p.X.Label.Text = "b, Impact Parameter"
	p.Y.Label.Text = "θ, Scattering Angle"

	xys := plotter.XYs{}
	for k, v := range bThetaMap {
		xys = append(xys, plotter.XY{X: k, Y: v})
	}
	bPlot, err := plotter.NewScatter(xys)
	if err != nil {
		panic(err)
	}
	bPlot.Shape = draw.CircleGlyph{}
	bPlot.Color = plotutil.Color(0)

	p.Add(bPlot)

	// Save the plot to a PNG file.
	if err := p.Save(5*vg.Inch, 5*vg.Inch, "scatter.png"); err != nil {
		panic(err)
	}

	os.Exit(1)
}
