package audioComponent

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/effects"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

func GetAudioFiles(directory string) []string {
	audioFiles := []string{}
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		audioFiles = append(audioFiles, file.Name())
	}
	return audioFiles
}

type audioPanel struct {
	sampleRate beep.SampleRate
	streamer   beep.StreamSeeker
	ctrl       *beep.Ctrl
	resampler  *beep.Resampler
	volume     *effects.Volume
}

func newAudioPanel(sampleRate beep.SampleRate, streamer beep.StreamSeeker) (audioPanel, error) {
	loopStreamer, err := beep.Loop2(streamer)
	if err != nil {
		return audioPanel{}, err
	}

	ctrl := &beep.Ctrl{Streamer: loopStreamer}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	return audioPanel{sampleRate, streamer, ctrl, resampler, volume}, nil
}

func (ap *audioPanel) play() {
	speaker.Play()
}
func (ap *audioPanel) stop() {
	speaker.Close()
}
func (ap *audioPanel) pause() {
	speaker.Lock()
	ap.ctrl.Paused = true
	speaker.Unlock()
}

func OpenAudioFile(audioFile string) (audioPanel, error) {
	audio, err := os.Open(audioFile)
	if err != nil {
		return audioPanel{}, err
	}
	streamer, format, err := mp3.Decode(audio)
	if err != nil {
		return audioPanel{}, err
	}
	ap, err := newAudioPanel(format.SampleRate, streamer)
	if err != nil {
		return audioPanel{}, err
	}
	return ap, nil
}

func PlayAudio(ap audioPanel) error {
	streamer := ap.streamer
	sampleRate := ap.sampleRate
	ap.play()
	duration := float64(streamer.Len()) / float64(sampleRate)

	time.Sleep(time.Duration(duration) * time.Second)
	ap.stop()
	return nil
}

func PauseAudio(ap audioPanel) error {
	sampleRate := ap.sampleRate
	speaker.Init(sampleRate, sampleRate.N(time.Second/30))
	ap.pause()
	return nil
}
