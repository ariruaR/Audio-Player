package audioComponent

import (
	"log"
	"os"
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
