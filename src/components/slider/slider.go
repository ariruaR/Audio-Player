package slider

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type SliderComponent struct {
	Container   *fyne.Container
	Slider      *widget.Slider
	CurrentTime *widget.Label
	TotalTime   *widget.Label
}

func NewSliderComponent(min, max float64, totalTime string, onChanged func(value float64)) *SliderComponent {
	currentTime := widget.NewLabel("0:00")
	totalTimeLabel := widget.NewLabel(totalTime)
	slider := widget.NewSlider(min, max)
	slider.Step = 1
	slider.OnChanged = onChanged

	timeRow := container.NewHBox(currentTime, layout.NewSpacer(), totalTimeLabel)

	c := container.NewVBox(slider, timeRow)
	return &SliderComponent{
		Container:   c,
		Slider:      slider,
		CurrentTime: currentTime,
		TotalTime:   totalTimeLabel,
	}
}
