package algo_prob

import (
	"math"
	"sort"
)

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func FMin(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

// BruteLine in a one dimension line places some discrete point,we want to find min dist point pair
// O(N*logN + N)
func BruteLine(points []int) int {
	sort.Ints(points)
	s, min := points[0], 0xffffff
	for i := 1; i < len(points); i++ {
		min = Min(min, points[i]-s)
		s = points[i]
	}
	return min
}

func Rec(points []int, l, r int) int {
	if r == l+1 {
		return points[r] - points[l]
	}
	mid := (r-l)/2 + l
	d1 := Rec(points, l, mid)
	d2 := Rec(points, mid, r)
	d := Min(d1, d2)
	d = Min(d, points[mid+1]-points[mid-1])
	return d
}

type Point struct {
	X, Y float64
}

func (p *Point) dist(f *Point) float64 {
	return math.Sqrt((p.X-f.X)*(p.X-f.X) + (p.Y-f.Y)*(p.Y-f.Y))
}

func Brute2Dimension(points []Point) float64 {
	min := 0xffffff * 1.00
	for i := 0; i < len(points); i++ {
		for j := i + 1; i < len(points); j++ {
			min = FMin(min, points[i].dist(&points[j]))
		}
	}
	return min
}
