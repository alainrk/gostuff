// Thanks to https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package golangz

import (
	"fmt"
)

type Vehicle interface {
	Steer(string) bool
	Accelerate(int) int
}

type Bike struct {
	rotationFreq float32
}
type Car struct {
	pedalDelta int
}

func (b *Bike) Steer (direction string) bool {
	b.rotationFreq += 12.341
	return true
}
func (c *Car) Steer (direction string) bool {
	c.pedalDelta += 42
	return true
}

func TestModels() {
	b := Bike{66.34}
	c := Car{12}
	fmt.Println(b, c)

	b.Steer("right")
	c.Steer("left")
	fmt.Println(b, c)
}

func Test() {
	TestModels()
}