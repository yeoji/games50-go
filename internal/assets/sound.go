package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
)

const SAMPLE_RATE = 44100

// map of sound name and sound file
type SoundLoaderConfig map[string][]byte

func loadSounds(audioContext *audio.Context, soundsToLoad SoundLoaderConfig) map[string]*audio.Player {
	sounds := make(map[string]*audio.Player)

	for name, sound := range soundsToLoad {
		sounds[name], _ = audio.NewPlayerFromBytes(audioContext, sound)
	}

	return sounds
}

func NewLoopingAudio(sound []byte) *audio.Player {
	audioContext := getAudioContext()

	stream, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(sound))
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
