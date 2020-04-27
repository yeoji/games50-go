package assets

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

const SAMPLE_RATE = 44100

// map of sound name and sound file
type SoundLoaderConfig map[string]string

func loadSounds(audioContext *audio.Context, soundsToLoad SoundLoaderConfig) map[string]*audio.Player {
	sounds := make(map[string]*audio.Player)

	for name, file := range soundsToLoad {
		sound, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Error reading sounds asset: %v", err)
		}
		sounds[name], _ = audio.NewPlayerFromBytes(audioContext, sound)
	}

	return sounds
}

func NewLoopingAudio(filePath string) *audio.Player {
	audioContext := getAudioContext()

	audioFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Could not open audio file: %v", err)
	}
	stream, err := mp3.Decode(audioContext, audioFile)
	if err != nil {
		log.Fatalf("Error decoding audio file: %v", err)
	}

	audioPlayer, err := audio.NewPlayer(audioContext, audio.NewInfiniteLoop(stream, stream.Length()))
	if err != nil {
		log.Fatalf("Error creating looping audio player: %v", err)
	}
	return audioPlayer
}

func getAudioContext() *audio.Context {
	if audio.CurrentContext() != nil {
		return audio.CurrentContext()
	}

	audioContext, err := audio.NewContext(SAMPLE_RATE)
	if err != nil {
		log.Fatalf("Could not initialize audio context: %v", err)
	}
	return audioContext
}
