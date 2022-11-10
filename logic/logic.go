package logic

import "fmt"

type Direction = int

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Grid = [4][4]int

type Logic struct {
	grid Grid
}

func New() *Logic {
	grid := Grid{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	l := &Logic{
		grid,
	}

	l.generateNewPiece()

	return l
}

func (l *Logic) generateNewPiece() {
	if !l.hasEmptySpots() {
		panic("attempted to generate new piece with no available slots")
	}

	value := randomValue()
loop:
	c, r := randomPos()
	if l.grid[c][r] != 0 {
		goto loop
	}
	l.grid[c][r] = int(value)
}

func (l *Logic) IsGameOver() bool {
	return l.hasLegalMoves() || l.hasEmptySpots()
}

func (l *Logic) hasLegalMoves() bool {
	for c := range l.grid {
		for r := range l.grid[c] {
			if c-1 >= 0 && l.grid[c][r] == l.grid[c-1][r] {
				return true
			}
			if c+1 < 4 && l.grid[c][r] == l.grid[c+1][r] {
				return true
			}
			if r-1 >= 0 && l.grid[c][r] == l.grid[c][r-1] {
				return true
			}
			if r+1 < 4 && l.grid[c][r] == l.grid[c][r+1] {
				return true
			}
		}
	}
	return false
}

func (l *Logic) hasEmptySpots() bool {
	for c := range l.grid {
		for r := range l.grid[c] {
			if l.grid[c][r] == 0 {
				return true
			}
		}
	}
	return false
}

func (l *Logic) moveToEmpty(ca int, ra int) {
outerLoop:
	hasMoved := false
	for c := range l.grid {
		for r := range l.grid[c] {
			if l.grid[c][r] != 0 {
				iter_c := 0
				incr_c := ca
				iter_r := 0
				incr_r := ra
			loop:
				if ca != 0 {
					if c+iter_c+incr_c >= 0 && c+iter_c+incr_c < 4 && l.grid[c+iter_c+incr_c][r] == 0 {
						l.grid[c+iter_c+incr_c][r] = l.grid[c+iter_c][r]
						l.grid[c+iter_c][r] = 0
						iter_c += incr_c
						hasMoved = true
						goto loop
					}
				}

				if ra != 0 {
					if r+iter_r+incr_r >= 0 && r+iter_r+incr_r < 4 && l.grid[c][r+iter_r+incr_r] == 0 {
						l.grid[c][r+iter_r+incr_r] = l.grid[c][r+iter_r]
						l.grid[c][r+iter_r] = 0
						iter_r += incr_r
						hasMoved = true
						goto loop
					}
				}
			}
		}
	}
	if hasMoved {
		goto outerLoop
	}
}

func (l *Logic) moveGeneral(ca int, ra int) {
	for c := range l.grid {
		for r := range l.grid[c] {
			if l.grid[c][r] != 0 {
				if ra != 0 && r+ra >= 0 && r+ra < 4 && l.grid[c][r] == l.grid[c][r+ra] {
					l.grid[c][r+ra] = l.grid[c][r] * 2
					l.grid[c][r] = 0
				} else if ca != 0 && c+ca >= 0 && c+ca < 4 && l.grid[c][r] == l.grid[c+ca][r] {
					l.grid[c+ca][r] = l.grid[c][r] * 2
					l.grid[c][r] = 0
				}
			}
		}
	}
	l.moveToEmpty(ca, ra)
}

func (l *Logic) MoveAndGenerate(direction Direction) {
	switch direction {
	case DirectionUp:
		l.moveGeneral(-1, 0)
	case DirectionDown:
		l.moveGeneral(1, 0)
	case DirectionRight:
		l.moveGeneral(0, 1)
	case DirectionLeft:
		l.moveGeneral(0, -1)
	}
	if l.hasEmptySpots() {
		l.generateNewPiece()
	}
	l.Display()
}

func (l *Logic) Data() Grid {
	return l.grid
}

func (l *Logic) Display() {
	for c := range l.grid {
		for r := range l.grid[c] {
			fmt.Printf("%d ", l.grid[c][r])
		}
		fmt.Print("\n")
	}
}
