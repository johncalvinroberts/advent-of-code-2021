package vector

import "github.com/johncalvinroberts/advent-of-code-2021/utils"

type XY struct{ X, Y int }

func (p1 XY) Subtract(p2 XY) XY {
	return p1.Add(p2.Multiply(-1))
}

func (p XY) Multiply(n int) XY {
	return XY{
		X: n * p.X,
		Y: n * p.Y,
	}
}

func (p XY) Sign() XY {
	return XY{
		X: utils.Sign(p.X),
		Y: utils.Sign(p.Y),
	}
}

func (p XY) Add(p2 XY) XY {
	return XY{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}
