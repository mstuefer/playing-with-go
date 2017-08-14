// OOP in go

package main

import (
	"fmt"
	"math"
)

// our object
type Point struct{ X, Y float64 }

// ad a so called receiver before the method-name
// receiver coz of "sending a message to an object"
func (p Point) EuclideanDistance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	fmt.Println("Point p(1,2)")
	p := Point{1, 2}
	fmt.Println("Point q(4,6)")
	q := Point{4, 6}
	fmt.Println("Distance: ", p.EuclideanDistance(q))
}
