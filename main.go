package main

import (
	"fmt"
	"time"

	termbox "github.com/nsf/termbox-go"
)

var (
	pos       int
	text      string
	startTime = time.Now()
	missCount int
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

	go clock()

	for _, v := range texts {
		text = v
		pos = 0
		for {
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
				} else {
					missCount++
				}
			}
		}
	}
}

func clock() {
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		// タイピング対象のテキスト
		setText(1, 1, text, pos)

		// 経過時間
		now := time.Now()
		diff := now.Sub(startTime)
		s := diff.Seconds()
		t := fmt.Sprintf("%.6f", s)
		setText(1, 3, t, 0)

		// 打ち間違いの数
		mt := fmt.Sprintf("%d", missCount)
		setText(1, 5, mt, 0)

		termbox.Flush()

		// 秒間60frame
		time.Sleep(1 * time.Second / 60)
	}
}

func setText(x, y int, text string, pos int) {
	dc := termbox.ColorDefault
	for i, ch := range text {
		c := dc
		if i < pos {
			c = termbox.ColorRed
		}
		termbox.SetCell(x+i, y, ch, c, dc)
	}
}
