package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	swapBtn := widget.NewButton("Swap", func() {
		l := left.Text
		left.SetText(right.Text)
		right.SetText(l)
	})

	openLeftBtn := widget.NewButton("Open←", func() {
		openFile(w, func(b []byte) { left.SetText(string(b)) })
	})
	openRightBtn := widget.NewButton("Open→", func() {
		openFile(w, func(b []byte) { left.SetText(string(b)) })
	})
	saveBtn := widget.NewButton("Save dif", func() {
		saveDiff(w, grid)
	})
	topBar := container.NewHBox(
		openLeftBtn, openRightBtn,
		widget.NewSeparator(),
		swapBtn, clearBtn,
		widget.NewSeparator(),
		diffBtn, saveBtn,
	)

}
