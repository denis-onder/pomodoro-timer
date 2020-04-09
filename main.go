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

func playNotifierSound() {
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

func sendNotification(arg string) {
	cmd := exec.Command("notify-send", arg)
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Running notify-send failed: %v", err)
		return
	}
	playNotifierSound()
}

func startPomodoro() {
	// Send out a notification, that the pomodoro timer started using notify-send
	sendNotification(`Pomodoro timer has started!` + "\n" + `Length: 25 minutes!`)
	// Sleep for 25 minutes
	time.Sleep(time.Minute * 25) // Real timer, uncomment before building
	// time.Sleep(time.Second * 5)
	// Notify the user that the timer has stopped
	sendNotification(`Timer stopped!` + "\n" + `Break period: 5 minutes!`)
	// Sleep for 5 minutes
	time.Sleep(time.Minute * 5) // Real timer, uncomment before building
	// time.Sleep(time.Second * 5)
	sendNotification(`Break ended!` + "\n" + `Starting new pomodoro cycle!`)
}

func handler() {
	flag := os.Args[1]
	switch flag {
	case "start":
		for {
			startPomodoro()
		}
	default:
		fmt.Println("Invalid flag!")
		break
	}
}

func main() {
	handler()
}
