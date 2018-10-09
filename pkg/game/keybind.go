package game

import (
	"github.com/k0kubun/termbox-go"
)

func WaitKeyInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				return
			}
		}
	}
}
