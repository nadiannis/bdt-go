package main

func main() {
	rectangle := Rectangle{
		Width:  5,
		Height: 7,
	}
	circle := Circle{
		Radius: 4,
	}

	PrintArea(rectangle)
	PrintArea(circle)
}
