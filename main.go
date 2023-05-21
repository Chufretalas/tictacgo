package main

func main() {
	TwoBotsPlayAGame()
}

func TwoBotsPlayAGame() {
	b := NewBoard()
	p1, ok1 := NewPlayer(&b, "O")
	p2, ok2 := NewPlayer(&b, "X")

	if ok1 == nil && ok2 == nil {
		for {
			if e := p1.PlayRandom(); e != nil {
				break
			}
			if s := b.Check(); s != "none" {
				break
			}
			if e := p2.PlayRandom(); e != nil {
				break
			}
			if s := b.Check(); s != "none" {
				break
			}
		}
	}
	b.Show()
}
