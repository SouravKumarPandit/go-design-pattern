package main

import (
	"fmt"
	"strings"
)

type MediaPlayer interface {
	play(audioType string, fileName string)
}

type AdvancedMediaPlayer interface {
	playVLC(fileName string)
	playMp4(fileName string)
}

type VlcPlayer struct{}

func (a VlcPlayer) playVLC(fileName string) {
	fmt.Printf("Playing VLC file. Name %v\n", fileName)
}

func (a VlcPlayer) playMp4(fileName string) {

}

type Mp4Player struct{}

func (a Mp4Player) playVLC(fileName string) {
}

func (a Mp4Player) playMp4(fileName string) {
	fmt.Printf("Playing Mp4 file. Name %v\n", fileName)

}

type mediaAdapter struct {
	mediaAdapter,
	amp AdvancedMediaPlayer
}

func (adapter *mediaAdapter) play(audioType string, fileName string) {

	if strings.EqualFold(audioType, "vlc") {
		adapter.amp.playVLC(fileName)
	} else
	if strings.EqualFold(audioType, "mp4") {
		adapter.amp.playMp4(fileName)

	}
}
func NewAdvanceMediaAdapter(audioType string) *mediaAdapter {
	if strings.EqualFold(audioType, "vlc") {
		return &mediaAdapter{amp: &VlcPlayer{}}

	} else
	if strings.EqualFold(audioType, "mp4") {
		return &mediaAdapter{amp: &Mp4Player{}}
	} else {
		return nil
	}
}

type audioPlayer struct {
	adapter mediaAdapter
}

func NewAudioPlayer() audioPlayer {
	return audioPlayer{
		adapter: mediaAdapter{},
	}
}

func (player audioPlayer) play(audioType string, fileName string) {
	amp := NewAdvanceMediaAdapter(audioType)
	if strings.EqualFold(audioType, "mp3") {
		fmt.Println("Playing mp3 file. Name : " + fileName)
	} else
	if strings.EqualFold(audioType, "mp4") || strings.EqualFold(audioType, "vlc") {
		amp.play(audioType, fileName)

	} else {
		fmt.Printf("Invalid media. . Name : %v ", fileName)

	}
}
