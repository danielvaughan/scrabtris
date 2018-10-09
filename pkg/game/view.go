package game

import (
	"github.com/nsf/termbox-go"
	"strings"
)

const (
	background = `
		WWWWWWWWWWWW WWWWWW
		WkkkkkkkkkkW WkkkkW
		WkkkkkkkkkkW WkkkkW
		WkkkkkkkkkkW WkkkkW
		WkkkkkkkkkkW WkkkkW
		WkkkkkkkkkkW WWWWWW
		WkkkkkkkkkkW
		WkkkkkkkkkkW
		WkkkkkkkkkkW BBBBBB
		WkkkkkkkkkkW WWWWWW
		WkkkkkkkkkkW
		WkkkkkkkkkkW
		WkkkkkkkkkkW 
		WkkkkkkkkkkW 
		WkkkkkkkkkkW
		WkkkkkkkkkkW 
		WkkkkkkkkkkW 
		WkkkkkkkkkkW
		WkkkkkkkkkkW
		WWWWWWWWWWWW

		kkkkkkkkkkkkkkkkkkk
		WWWWWWWWWWWWWWWWWWW
	`
	boardXOffset, boardYOffset = 3, 2
)

var (
	colorMapping = map[rune]termbox.Attribute{
		'k': termbox.ColorBlack,
		'K': termbox.ColorBlack | termbox.AttrBold,
		'W': termbox.ColorWhite | termbox.AttrBold,
	}
)

func RewriteScreen(text string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawBacks(background, 0, 0)
	drawCells(text, boardXOffset, boardYOffset)
	drawTexts()

	termbox.Flush()
}

func drawTexts() {
	drawText(32, 9, "SCORE", termbox.ColorWhite, termbox.ColorBlue)
	drawText(3, 22, "  ←           <SPC>          →", termbox.ColorWhite, termbox.ColorBlack)
	drawText(3, 23, " left         drop         right", termbox.ColorBlack, termbox.ColorWhite)
	drawText(30, 20, " q: quit", termbox.ColorWhite, termbox.ColorDefault)
}

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for index, ch := range text {
		termbox.SetCell(x+index, y, rune(ch), fg, bg)
	}
}

func drawCells(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			drawText(left+x, top+y, string(char), termbox.ColorWhite, termbox.ColorDefault)
		}
	}
}

func drawBacks(text string, left, top int) {
	lines := strings.Split(text, "\n")

	for y, line := range lines {
		for x, char := range line {
			drawBack(left+x, top+y, colorByChar(char))
		}
	}
}

func drawBack(x, y int, color termbox.Attribute) {
	termbox.SetCell(2*x-1, y, ' ', termbox.ColorDefault, color)
	termbox.SetCell(2*x, y, ' ', termbox.ColorDefault, color)
}

func colorByChar(ch rune) termbox.Attribute {
	return colorMapping[ch]
}

func charByColor(color termbox.Attribute) rune {
	for ch, currentColor := range colorMapping {
		if currentColor == color {
			return ch
		}
	}
	return '.'
}
