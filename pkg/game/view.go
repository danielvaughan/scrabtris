package game

import (
	"github.com/danielvaughan/scrabtris/pkg/board"
	"github.com/nsf/termbox-go"
	"strings"
)

type View struct {
	Board *board.Board
	Game  *Game
}

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
	boardXOffset, boardYOffset = 5, 2
)

var (
	colorMapping = map[rune]termbox.Attribute{
		'k': termbox.ColorBlack,
		'B': termbox.ColorBlue | termbox.AttrBold,
		'W': termbox.ColorWhite | termbox.AttrBold,
	}
)

func (v *View) refreshScreen() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawBacks(background, 0, 0)
	drawCells(v.Board.State(), boardXOffset, boardYOffset)
	drawNextTile(v.Game.nextTile.Letter)
	drawTexts()

	termbox.Flush()
}

func drawNextTile(letter rune) {
	drawText(34, 3, string(letter), termbox.ColorBlack, termbox.ColorYellow)
}

func drawTexts() {
	drawText(32, 9, "SCORE", termbox.ColorWhite, termbox.ColorBlue)
	drawText(3, 22, " ←                            →", termbox.ColorWhite, termbox.ColorBlack)
	drawText(3, 23, " left                     right", termbox.ColorBlack, termbox.ColorWhite)
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
			if char == ' ' {
				drawText(left+2*x, top+y, string(char), termbox.ColorBlack, termbox.ColorBlack)
			} else {
				drawText(left+2*x, top+y, string(char), termbox.ColorBlack, termbox.ColorYellow)
			}
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
