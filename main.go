package main

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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

	//レイアウト
	inputs := container.NewHSplit(
		container.NewBorder(nil, nil, nil, nil, left),
		container.NewBorder(nil, nil, nil, nil, right),
	)
	inputs.SetOffset(0.5)

	w.SetContent(container.NewBorder(
		topBar, nil, nil, nil,
		container.NewVSplit(
			inputs,
			container.NewMax(grid),
		),
	))
	w.ShowAndRun()
}

func openFile(w fyne.Window, onLoad func([]byte)) {
	d := dialog.NewFileOpen(func(rc fyne.URIReadCloser, err error) {
		if err != nil || rc == nil {
			return
		}
		defer rc.Close()
		b, _ := io.ReadAll(rc)
		onLoad(b)
	}, w)
	d.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".log", ".md", ".json", ".yaml", ".yml", ""}))
	d.Show()
}
