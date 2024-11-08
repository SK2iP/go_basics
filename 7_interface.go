package main
//import fmt and math
import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}
// getArea function takes a shape interface and returns the area of the shape
func getArea(s shape) float64 {
	sh , ok := s.(rect)
	if ok {
		fmt.Printf("Shape is a rectangle with width %f and height %f\n", sh.width, sh.height)
		return sh.area()
	}
	new_sh , ok := s.(circle)
	if ok {
		fmt.Printf("Shape is a circle with radius %f\n", new_sh.radius)
		return sh.area()
	}
	return 0
}

func main() {
	r := rect{width: 12, height: 9}
	c := circle{radius: 6}
	measure(r)
	measure(c)

	fmt.Println(getArea(r))
}


