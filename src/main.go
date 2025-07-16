package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500, 500))
	title := widget.NewLabelWithStyle("title", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: false,
		Underline: false,
	})
	artist := widget.NewLabelWithStyle("artist", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: false,
		Underline: false,
	})
	rectangle := canvas.NewRectangle(color.White)
	rectangle.SetMinSize(fyne.NewSize(100, 100))

	musicCover := container.New(layout.NewCenterLayout(), rectangle)

	progressbar := widget.NewProgressBar()
	progressbar.SetValue(0.25)
	progressbar.Resize(fyne.NewSize(200, 40))
	sizedBar := container.New(layout.NewCenterLayout(), progressbar)
	sizedBar.Resize(fyne.NewSize(30, 300))
	btnLabel := "PLAY"
	w.SetContent(container.NewVBox(
		musicCover,
		title,
		artist,
		sizedBar,
		widget.NewButton(btnLabel, func() {}),
	))

	w.ShowAndRun()
}
