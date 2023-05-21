package main

func main() {
	b := NewBoard()
	b.Show()
	b.X(0, 2)
	b.X(1, 1)
	b.X(2, 0)
	b.Show()
	b.Check()
}
