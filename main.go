package main

import (
	termbox "github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	texts := []string{
		"golang",
		"java",
		"clang",
		"objectivec",
	}

	for _, text := range texts {
		var pos int
		for {
			setText(text, pos)
			if len(text) <= pos {
				break
			}
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyCtrlC, termbox.KeyCtrlD:
					return
				}
				ch := ev.Ch
				top := rune(text[pos])
				same := ch == top
				if same {
					pos++
				}
			}
		}
	}

}

func setText(text string, pos int) {
	dc := termbox.ColorDefault
	termbox.Clear(dc, dc)
	for i, ch := range text {
		c := dc
		if i < pos {
			c = termbox.ColorRed
		}
		termbox.SetCell(1+i, 1, ch, c, dc)
	}
	termbox.Flush()
}
