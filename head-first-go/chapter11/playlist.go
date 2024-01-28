package main

import "github.com/nbompetsis/gopackages/gadget"

func playList(t gadget.Player, songs []string) {
	for _, song := range songs {
		t.Play(song)
	}
	t.Stop()
}

func main() {
	var device gadget.Player
	// device := gadget.TapeRecorder{}
	device = gadget.TapePlayer{}
	songs := []string{"Jessie's Girl", "Whip It"}
	songs = append(songs, "9 to 5")
	playList(device, songs)

	device = gadget.TapeRecorder{}
	songs = []string{"Walkman"}
	playList(device, songs)
}
