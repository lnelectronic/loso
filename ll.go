// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 29/3/2564 22:14
// ---------------------------------------------------------------------------
package main

import (
	"fmt"
	"math"
	"reflect"
)

type Joy interface {
	area() float32
}

type rec struct {
	width  float32
	height float32
	test   int
}
type eiei struct {
	radius float32
}

func (c eiei) area() float32 {
	return math.Pi * c.radius * c.radius

}

func (r rec) area() float32 {
	return r.width * r.height
}

func getArea(s Joy) float32 {
	return s.area()
}

func checkType(s Joy) {
	t := reflect.TypeOf(s).Name()
	switch t {
	case "rec":
		r := s.(rec)
		fmt.Printf("Rec  width: %f, height: %f\n", r.width, r.height)
		fmt.Println("test:", r.test)
		break
	case "eiei":
		c := s.(eiei)
		fmt.Printf("Circle radius: %f\n", c.radius)
	}

}

func main() {

	l := rec{width: 20, height: 20, test: 100}
	rr := eiei{radius: 15}

	fmt.Println("Rectangle: %f", getArea(l))
	fmt.Println("Circle: %f", getArea(rr))
	fmt.Println("-------------------------")
	checkType(l)

}
