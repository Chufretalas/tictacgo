package main

import "fmt"

func main() {
	TestRandomOdds(1000)
}

// Teste the odds of one random bot winning againsty another
func TestRandomOdds(turns int) {
	results := make(map[string]int)
	results["X"] = 0
	results["O"] = 0
	results["none"] = 0
	for i := 0; i < turns; i++ {
		results[TwoBotsPlayAGame(false)] += 1
	}
	fmt.Println(results)
}

func TwoBotsPlayAGame(printResult bool) string {
	b := NewBoard()
	p1, ok1 := NewPlayer(&b, "O")
	p2, ok2 := NewPlayer(&b, "X")

	if ok1 == nil && ok2 == nil {
		for {
			if e := p1.PlayRandom(); e != nil {
				break
			}
			if s := b.Check(printResult); s != "none" {
				break
			}
			if e := p2.PlayRandom(); e != nil {
				break
			}
			if s := b.Check(printResult); s != "none" {
				break
			}
		}
	}
	if printResult {
		b.Show()
	}
	return b.Check(false)
}
