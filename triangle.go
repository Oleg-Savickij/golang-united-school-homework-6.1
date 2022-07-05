package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (r Triangle) CalcPerimeter() float64 {
	return r.Side * 3
}

func (r Triangle) CalcArea() float64 {
	return (math.Pow(r.Side, 2) * math.Sqrt(3)) / 4
}
