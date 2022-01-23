package main

import "fmt"

type Test struct {
	ID        int
	AreaValue int
	AreaType  string
}

func main() {
	var a Test
	_ = a.New(10, 10, "persegi")
	fmt.Println(a)
}

func (ar *Test) New(param1, param2 int, shape string) (err error) {
	ar.AreaType = shape
	switch shape {
	case "persegi panjang":
		ar.AreaValue = param1 * param2
	case "persegi":
		ar.AreaValue = param1 * param2
	case "segitiga":
		ar.AreaValue = (param1 * param2) / 2
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
	}

	return nil
}
