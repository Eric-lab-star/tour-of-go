package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v Vertex) add() float64 {
	return v.X + v.Y
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func mutate(v *Vertex) {
	v.X = v.X * 2
	v.Y = v.Y * 2
}

func main() {
	v := Vertex{3, 2}
	fmt.Println(v)
	mutate(&v)
	fmt.Println(v)

}
