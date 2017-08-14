// OOP in go

package main

import (
	"fmt"
	"image/color"
	"math"
)

// our object
type Point struct{ X, Y float64 }

// ad a so called receiver before the method-name
// receiver coz of "sending a message to an object"
func (p Point) EuclideanDistance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// usually calling a function make a copy of each argument value, whenever
// we need to update a variable, or an argument is too large and we therefore
// whish to avoid copying, we pass the address of the variable using a pointer
// the same goes for methods that need to update the receiver variable
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// Composing Types by Struct Embedding
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func main() {
	fmt.Println("Point p(1,2)")
	p := Point{1, 2}
	fmt.Println("Point q(4,6)")
	q := Point{4, 6}
	fmt.Println("Distance: ", p.EuclideanDistance(q))

	fmt.Println("Scale point p by 2")
	p.ScaleBy(2) // should be (&p).ScaleBy .. however, the compiler helps us out, translating p -> (&p)
	// If the receiver p is a variable of type Point but the method requires *Point receiver, compiler will perform implicit &p on the variable
	fmt.Println("p.X:", p.X, " p.Y:", p.Y)

	var cp ColoredPoint
	cp.X = 100 // embedding let us use X and Y directly without writing xp.Point.X, although that works as well
	cp.Y = 200
	cp.Color = color.RGBA{255, 0, 0, 255}
	fmt.Println("cp.X:", cp.X, " cp.Y:", cp.Y, " cp.Color:", cp.Color)
}
