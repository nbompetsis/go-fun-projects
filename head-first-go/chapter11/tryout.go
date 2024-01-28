package main

import (
	"fmt"

	"github.com/nbompetsis/gopackages/gadget"
)

func TryOut(p gadget.Player) {
	p.Play("tryout-song")
	p.Stop()

	tapePlayer, ok := p.(gadget.TapePlayer)
	if ok {
		fmt.Println("Gadget betteries:", tapePlayer.Batteries)
	} else {
		fmt.Println("The original type is not TapePlayer. The assertion was unsuccessful")
	}
}

func main() {
	TryOut(gadget.TapePlayer{Batteries: "AAA+"})
	TryOut(gadget.TapeRecorder{Microphone: 1})

}
