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
		for {
			setText(text)
			if len(text) < 1 {
				break
			}
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyCtrlC, termbox.KeyCtrlD:
					return
				}
				ch := ev.Ch
				top := rune(text[0])
				same := ch == top
				if same {
					text = text[1:]
				}
			}
		}
	}

}

func setText(text string) {
	dc := termbox.ColorDefault
	termbox.Clear(dc, dc)
	for i, ch := range text {
		termbox.SetCell(1+i, 1, ch, dc, dc)
	}
	termbox.Flush()
}
