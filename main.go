package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playNotifier() {
	f, err := os.Open("ding.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func startPomodoro() {
	// Send out a notification, that the pomodoro timer started using notify-send
	cmd := exec.Command("notify-send", `Pomodoro timer has started!`+"\n"+`Length: 25 Minutes!`)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Running notify-send failed: %v", err)
		return
	}
	// Play notifier sound
	playNotifier()
	// Sleep for 25 minutes
	// time.Sleep(time.Minute * 25)
	time.Sleep(time.Second * 5)
	// Notify the user that the timer has been
	return
}

func handler() {
	flag := os.Args[1]
	switch flag {
	case "start":
		startPomodoro()
		break
	case "stop":
		fmt.Println("stop")
		break
	default:
		fmt.Println("invalid flag")
		break
	}
}

func main() {
	handler()
}
