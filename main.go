package main

import (
	"image/color"
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pmezard/go-difflib/difflib"
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

func doDiff(left, right string, grid *widget.TextGrid) {
	ud := difflib.UnifiedDiff{
		A:        splitLines(left),
		B:        splitLines(right),
		FromFile: "left",
		ToFile:   "right",
		Context:  3,
	}
	out, _ := difflib.GetUnifiedDiffString(ud)
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")

	grid.SetText("")
	grid.Rows = nil

	// 色定義
	red := color.NRGBA{R: 220, G: 72, B: 72, A: 255} // 追加/削除強調色
	green := color.NRGBA{R: 46, G: 204, B: 113, A: 255}
	gray := color.NRGBA{R: 140, G: 140, B: 140, A: 255}
	cyan := color.NRGBA{R: 52, G: 152, B: 219, A: 255}
	def := color.NRGBA{R: 230, G: 230, B: 230, A: 255} // デフォルト/薄め

	for _, ln := range lines {
		row := widget.TextGridRow{}
		var fg color.Color = def

		switch {
		case strings.HasPrefix(ln, "@@"):
			fg = cyan
		case strings.HasPrefix(ln, "---") || strings.HasPrefix(ln, "+++"):
			fg = gray
		case strings.HasPrefix(ln, "-"):
			fg = red
		case strings.HasPrefix(ln, "+"):
			fg = green
		default:
			fg = def
		}

		cells := make([]widget.TextGridCell, len(ln))
		for i, ch := range ln {
			cells[i] = widget.TextGridCell{
				Rune:  ch,
				Style: widget.TextGridStyle{FGColor: fg},
			}
		}
		row.Cells = cells
		grid.Rows = append(grid.Rows, row)
	}
	grid.Refresh()
}
