package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	audioComponent "AudioPlayer/src/components/audioComponent"
	slider "AudioPlayer/src/components/slider"
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

	musicData := audioComponent.GetAudioFiles("C:/Users/Ariruar/Desktop/AudioPlayer/Audio-Player/src/audio")

	currentSong := 0

	updateSong := func() {
		title.Text = musicData[currentSong]
		title.Refresh()
		// Здесь можно обновлять artist.Text, если есть данные об исполнителе
	}

	rectangle := canvas.NewRectangle(color.White)
	rectangle.SetMinSize(fyne.NewSize(100, 100))
	shadow := canvas.NewRectangle(color.RGBA{255, 248, 220, 255})
	shadow.Resize(fyne.NewSize(110, 110))
	shadowContainer := container.NewStack(
		shadow,
		rectangle,
	)
	sliderComponent := slider.NewSliderComponent(0, 180, "3:00", func(value float64) {})
	sliderComponent.Slider.OnChanged = func(value float64) {
		sliderComponent.CurrentTime.SetText(fmt.Sprintf("%d:%d", int(value/60), int(value)%60))
	}

	musicCover := container.New(layout.NewCenterLayout(), shadowContainer)

	playBtn := widget.NewButton("PLAY", func() {})
	backBtn := widget.NewButton("<=", func() {
		if currentSong > 0 {
			currentSong--
			updateSong()
			sliderComponent.Slider.Value = 0
			sliderComponent.CurrentTime.SetText("0:00")
			sliderComponent.CurrentTime.Refresh()
			sliderComponent.Slider.Refresh()
		}
	})
	nextBtn := widget.NewButton("=>", func() {
		if currentSong < len(musicData)-1 {
			currentSong++
			updateSong()
			sliderComponent.Slider.Value = 0
			sliderComponent.CurrentTime.SetText("0:00")
			sliderComponent.CurrentTime.Refresh()
			sliderComponent.Slider.Refresh()
		}
	})
	playBtn.Resize(fyne.NewSize(150, 50))
	btnContainer := container.NewCenter(
		container.NewHBox(backBtn, playBtn, nextBtn),
	)
	background := canvas.NewRectangle(color.RGBA{50, 50, 50, 255})

	content := container.NewVBox(
		musicCover,
		title,
		artist,
		sliderComponent.Container,
		btnContainer,
	)

	musicList := widget.NewList(
		func() int { return len(musicData) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(musicData[i])
		},
	)
	musicList.Resize(fyne.NewSize(100, musicList.Size().Height))
	musicList.OnSelected = func(id widget.ListItemID) {
		currentSong = id
		updateSong()
	}
	contentContainer := container.NewGridWithColumns(2, musicList, container.NewCenter(content))
	mainContainer := container.NewStack(
		background,
		contentContainer,
	)

	w.SetContent(mainContainer)
	updateSong()

	w.ShowAndRun()
}
