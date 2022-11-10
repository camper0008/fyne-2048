package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/camper0008/fyne-2048/logic"
)

func updateBindings(data logic.ViewGrid, bindings *[]binding.String) {
	for c := range data {
		for r := range data[c] {
			_ = (*bindings)[c*4+r].Set(data[c][r])
		}
	}
}

func generateBoundContainer(data logic.ViewGrid, window fyne.Window) (*fyne.Container, *[]binding.String) {
	bindings := make([]binding.String, 16)
	objects := make([]fyne.CanvasObject, 16)
	for c := range data {
		for r := range data[c] {
			bindings[c*4+r] = binding.NewString()
			_ = bindings[c*4+r].Set(data[c][r])
			objects[c*4+r] = container.NewCenter(widget.NewLabelWithData(bindings[c*4+r]))
		}
	}

	return container.New(layout.NewGridLayout(4), objects...), &bindings
}

func main() {
	rand.Seed(time.Now().UnixNano())
	gameOverScreen := false
	l := logic.New()
	a := app.New()
	w := a.NewWindow("2048")
	c, b := generateBoundContainer(l.View(), w)
	scoreBinding := binding.NewString()
	w.SetContent(container.NewVBox(container.NewCenter(widget.NewLabelWithData(scoreBinding)), c))
	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		if l.IsGameOver() {
			if !gameOverScreen {
				dialog.ShowInformation("Game over!", "du dårlig", w)
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
		updateBindings(l.View(), b)
		scoreBinding.Set(l.FormattedScore())
	})
	w.Resize(fyne.NewSize(w.Canvas().Size().Width+25, w.Canvas().Size().Height))
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}
