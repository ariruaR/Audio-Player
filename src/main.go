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
	title := canvas.NewText("Title", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 25

	artist := canvas.NewText("Artist", color.White)
	artist.Alignment = fyne.TextAlignCenter
	artist.TextSize = 15

	rectangle := canvas.NewRectangle(color.White)
	rectangle.SetMinSize(fyne.NewSize(100, 100))
	shadow := canvas.NewRectangle(color.RGBA{255, 248, 220, 255})
	shadow.Resize(fyne.NewSize(110, 110))
	shadowContainer := container.NewStack(
		shadow,
		rectangle,
	)

	musicCover := container.New(layout.NewCenterLayout(), shadowContainer)

	playBtn := widget.NewButton("PLAY", func() {})
	backBtn := widget.NewButton("<=", func() {})
	nextBtn := widget.NewButton("=>", func() {})
	playBtn.Resize(fyne.NewSize(150, 50))
	btnContainer := container.NewCenter(
		container.NewHBox(backBtn, playBtn, nextBtn),
	)
	background := canvas.NewRectangle(color.RGBA{50, 50, 50, 255})

	slider := widget.NewSlider(0, 180) // 180 секунд — пример
	slider.Value = 0
	slider.Step = 1
	slider.OnChanged = func(value float64) {
		// Здесь можно обновлять отображение времени или перематывать трек
	}

	currentTime := widget.NewLabel("0:00")
	totalTime := widget.NewLabel("3:00")
	timeRow := container.NewHBox(currentTime, widget.NewLabel(" "), totalTime)

	content := container.NewVBox(
		musicCover,
		title,
		artist,
		slider,
		timeRow,
		btnContainer,
	)
	musicData := []string{
		"Song 1",
		"Song 2",
		"Song 3",
		"Song 4",
		"Song 5",
		"Song 6",
		"Song 7",
		"Song 8",
		"Song 9",
		"Song 10",
		"Song 11",
		"Song 12",
		"Song 13",
		"Song 14",
	}

	musicList := widget.NewList(
		func() int { return len(musicData) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(musicData[i])
		},
	)
	musicList.Resize(fyne.NewSize(100, musicList.Size().Height))
	musicList.OnSelected = func(id widget.ListItemID) {
		selectedSong := musicData[id]
		title.Text = selectedSong
	}
	contentContainer := container.NewGridWithColumns(2, musicList, container.NewCenter(content))
	// Накладываем контент на фо
	mainContainer := container.NewStack(
		background,
		contentContainer,
	)

	w.SetContent(mainContainer)

	w.ShowAndRun()
}
