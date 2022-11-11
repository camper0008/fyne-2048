package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/camper0008/fyne-2048/logic"
)

type cell struct {
	label      *canvas.Text
	background *canvas.Rectangle
}

func colorsFromCellValue(value int) (color.Color, color.Color) {
	if value == 0 {
		return color.RGBA{}, color.RGBA{}
	}
	if value <= 64 {
		x := math.Log2(float64(value))
		v := uint8(255 - x*(255/6))
		var textColor color.Color
		if v < 127 {
			textColor = color.RGBA{R: 1, G: 1, B: 1, A: 255}
		} else {
			textColor = color.RGBA{R: 200, G: 255, B: 255, A: 255}
		}
		return color.RGBA{R: 255, G: v, B: v}, textColor
	}
	return color.RGBA{}, color.RGBA{}
}

func updateCells(data logic.DataGrid, cells *[]cell) {
	for c := range data {
		for r := range data[c] {
			bgColor, txtColor := colorsFromCellValue(data[c][r])
			cell := (*cells)[c*4+r]
			cell.label.Text = fmt.Sprintf("%d", data[c][r])
			cell.label.Color = txtColor
			cell.background.FillColor = bgColor
			cell.label.Refresh()
			cell.background.Refresh()
		}
	}
}

func generateBoundContainer(data logic.ViewGrid, window fyne.Window) (*fyne.Container, *[]cell) {
	cells := make([]cell, 16)
	objects := make([]fyne.CanvasObject, 16)
	for c := range data {
		for r := range data[c] {
			label := &canvas.Text{}
			background := &canvas.Rectangle{}
			background.SetMinSize(fyne.Size{50, 50})
			cells[c*4+r] = cell{
				label,
				background,
			}
			objects[c*4+r] = container.NewCenter(label, background)
		}
	}

	return container.New(layout.NewGridLayout(4), objects...), &cells
}

func main() {
	rand.Seed(time.Now().UnixNano())
	gameOverScreen := false
	l := logic.New()
	a := app.New()
	w := a.NewWindow("2048")
	c, b := generateBoundContainer(l.View(), w)
	updateCells(l.Data(), b)
	scoreBinding := binding.NewString()
	scoreBinding.Set(l.FormattedScore())
	w.SetContent(container.NewVBox(container.NewCenter(widget.NewLabelWithData(scoreBinding)), c))
	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if l.IsGameOver() {
			if !gameOverScreen {
				dialog.ShowInformation("Game over!", "", w)
			}
			gameOverScreen = true
			return
		}
		switch key.Name {
		case fyne.KeyDown:
			l.MoveAndGenerate(logic.DirectionDown)
		case fyne.KeyUp:
			l.MoveAndGenerate(logic.DirectionUp)
		case fyne.KeyLeft:
			l.MoveAndGenerate(logic.DirectionLeft)
		case fyne.KeyRight:
			l.MoveAndGenerate(logic.DirectionRight)
		}
		updateCells(l.Data(), b)
		scoreBinding.Set(l.FormattedScore())
	})
	w.Resize(fyne.NewSize(w.Canvas().Size().Width, w.Canvas().Size().Height))
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}
