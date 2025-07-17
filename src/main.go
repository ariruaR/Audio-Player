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
	shadow := canvas.NewRectangle(color.RGBA{255, 248, 220, 255})
	shadow.Resize(fyne.NewSize(110, 110))
	shadowContainer := container.NewStack(
		shadow,
		rectangle,
	)

	musicCover := container.New(layout.NewCenterLayout(), shadowContainer)

	progressbar := widget.NewProgressBar()
	progressbar.SetValue(0.25)
	progressbar.Resize(fyne.NewSize(200, 40))
	sizedBar := container.New(layout.NewCenterLayout(), progressbar)
	sizedBar.Resize(fyne.NewSize(30, 300))
	playBtn := widget.NewButton("PLAY", func() {})
	backBtn := widget.NewButton("<=", func() {})
	nextBtn := widget.NewButton("=>", func() {})
	playBtn.Resize(fyne.NewSize(150, 50))
	btnContainer := container.NewCenter(
		container.NewHBox(backBtn, playBtn, nextBtn),
	)
	background := canvas.NewRectangle(color.RGBA{50, 50, 50, 255})

	content := container.NewVBox(
		musicCover,
		title,
		artist,
		sizedBar,
		btnContainer,
	)

	// Накладываем контент на фон
	mainContainer := container.NewStack(
		background,
		container.NewCenter(content),
	)

	w.SetContent(mainContainer)

	w.ShowAndRun()
}
