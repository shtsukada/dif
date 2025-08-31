package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("dif-fyne")
	w.Resize(fyne.NewSize(1000, 700))

	//入力フィールド
	left := widget.NewMultiLineEntry()
	left.SetPlaceHolder("left")
	right := widget.NewMultiLineEntry()
	right.SetPlaceHolder("right")

	grid := widget.NewTextGrid()
	grid.ShowLineNumbers = true
	grid.SetText("unified diff")

	diffBtn := widget.NewButton("dif", func() {
		doDiff(left.Text, right.Text, grid)
	})
	clearBtn := widget.NewButton("Clear", func() {
		left.SetText("")
		right.SetText("")
		grid.SetText("")
	})

}
