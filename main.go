package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/camper0008/fyne-2048/logic"
)

func updateBindings(data logic.Grid, bindings *[]binding.Int) {
	for c := range data {
		for r := range data[c] {
			_ = (*bindings)[c*4+r].Set(data[c][r])
		}
	}
}

func generateBoundContainer(data logic.Grid, window fyne.Window) (*fyne.Container, *[]binding.Int) {
	bindings := make([]binding.Int, 16)
	objects := make([]fyne.CanvasObject, 16)
	for c := range data {
		for r := range data[c] {
			bindings[c*4+r] = binding.NewInt()
			_ = bindings[c*4+r].Set(data[c][r])
			objects[c*4+r] = container.NewCenter(widget.NewLabelWithData(binding.IntToStringWithFormat(bindings[c*4+r], "%d")))
		}
	}

	return container.New(layout.NewGridLayout(4), objects...), &bindings
}

func main() {
	rand.Seed(time.Now().UnixNano())
	l := logic.New()
	a := app.New()
	w := a.NewWindow("2048")
	c, b := generateBoundContainer(l.Data(), w)
	w.SetContent(c)
	w.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
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
		updateBindings(l.Data(), b)
	})
	w.Resize(fyne.NewSize(w.Canvas().Size().Width+200, w.Canvas().Size().Height+200))
	w.CenterOnScreen()
	w.SetMaster()
	w.ShowAndRun()
}
