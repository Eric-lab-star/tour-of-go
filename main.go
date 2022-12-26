package main

type student struct {
	name string
	age  int
}

func main() {
	var p *student
	kim := student{
		name: "kim",
		age:  13,
	}
	p = &kim

	p.name = "Lee"

}
