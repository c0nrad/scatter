package main

type Vec struct {
	X, Y float64
}

func (v1 Vec) Add(v2 Vec) Vec {
	return Vec{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 Vec) Minus(v2 Vec) Vec {
	return Vec{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v1 Vec) Dot(v2 Vec) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

func (v1 Vec) Scale(a float64) Vec {
	return Vec{X: v1.X * a, Y: v1.Y * a}
}

func (u Vec) ProjectOnto(v Vec) Vec {
	return v.Scale(u.Dot(v) / v.Dot(v))
}

func (v1 Vec) DistanceSquared(v2 Vec) float64 {
	return v1.Minus(v2).Dot(v1.Minus(v2))
}
