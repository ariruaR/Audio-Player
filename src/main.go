package main

import (
	"fmt"
	"image/color"
	"time"

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

	musicData := []string{
		"Song 1", "Song 2", "Song 3", "Song 4", "Song 5",
		"Song 6", "Song 7", "Song 8", "Song 9", "Song 10",
		"Song 11", "Song 12", "Song 13", "Song 14",
	}

	currentSong := 0
	isPlaying := false
	var sliderTimer *time.Timer

	rectangle := canvas.NewRectangle(color.White)
	rectangle.SetMinSize(fyne.NewSize(100, 100))
	shadow := canvas.NewRectangle(color.RGBA{255, 248, 220, 255})
	shadow.Resize(fyne.NewSize(110, 110))
	shadowContainer := container.NewStack(shadow, rectangle)
	musicCover := container.New(layout.NewCenterLayout(), shadowContainer)

	slider := widget.NewSlider(0, 180)
	slider.Value = 0
	slider.Step = 1

	currentTime := widget.NewLabel("0:00")
	totalTime := widget.NewLabel("3:00")
	timeRow := container.NewHBox(currentTime, widget.NewLabel(" "), totalTime)

	// Функция для обновления отображения времени
	updateTimeDisplay := func(current, total float64) {
		currentMin := int(current) / 60
		currentSec := int(current) % 60
		totalMin := int(total) / 60
		totalSec := int(total) % 60
		currentTime.SetText(fmt.Sprintf("%d:%02d", currentMin, currentSec))
		totalTime.SetText(fmt.Sprintf("%d:%02d", totalMin, totalSec))
	}

	// Объявляем playBtn как переменную
	var playBtn *widget.Button

	// Функция для анимации слайдера
	animateSlider := func() {
		if sliderTimer != nil {
			sliderTimer.Stop()
		}
		sliderTimer = time.NewTimer(1 * time.Second)
		go func() {
			for range sliderTimer.C {
				if isPlaying && slider.Value < 180 {
					slider.Value++
					slider.Refresh()
					updateTimeDisplay(slider.Value, 180)
					sliderTimer = time.NewTimer(1 * time.Second)
				} else {
					isPlaying = false
					playBtn.SetText("PLAY")
					break
				}
			}
		}()
	}

	// Функция для обновления песни
	updateSong := func() {
		title.Text = musicData[currentSong]
		title.Refresh()
		// Сброс слайдера при смене песни
		slider.Value = 0
		slider.Refresh()
		updateTimeDisplay(0, 180)
	}

	// Настройка слайдера
	slider.OnChanged = func(value float64) {
		updateTimeDisplay(value, 180)
	}

	// Создание кнопок
	playBtn = widget.NewButton("PLAY", func() {
		isPlaying = !isPlaying
		if isPlaying {
			playBtn.SetText("PAUSE")
			animateSlider()
		} else {
			playBtn.SetText("PLAY")
			if sliderTimer != nil {
				sliderTimer.Stop()
			}
		}
	})

	backBtn := widget.NewButton("<=", func() {
		if currentSong > 0 {
			currentSong--
		} else {
			currentSong = len(musicData) - 1
		}
		updateSong()
	})

	nextBtn := widget.NewButton("=>", func() {
		if currentSong < len(musicData)-1 {
			currentSong++
		} else {
			currentSong = 0
		}
		updateSong()
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
		slider,
		timeRow,
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
	mainContainer := container.NewStack(background, contentContainer)

	w.SetContent(mainContainer)
	updateSong()

	w.ShowAndRun()
}
