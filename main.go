package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playNotifierSound() {
	home, _ := os.LookupEnv("HOME")
	f, err := os.Open(home + "/Music/ding.mp3")
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

func startPomodoro(cycles int64) {
	var minutesPassed int64 = 0
	// Send out a notification, that the pomodoro timer started using notify-send
	sendNotification("Pomodoro timer\nThe timer has started!\nCycles: " + strconv.FormatInt(cycles, 10) + "\n Length: 25 minutes!")
	// Sleep for 25 minutes
	for i := 0; i < 25; i++ {
		time.Sleep(time.Minute * 1) // Real timer
		// time.Sleep(time.Second * 1)
		minutesPassed++
		t := 25 - minutesPassed
		sendNotification("Pomodoro timer\nTime left: " + strconv.FormatInt(t, 10))
	}
	// Notify the user that the timer has stopped
	sendNotification("Pomodo timer\nCycle has ended.\nStarting a 5 minute break period!")
	// Sleep for 5 minutes
	minutesPassed = 0
	for i := 0; i < 5; i++ {
		time.Sleep(time.Minute * 1) // Real timer
		// time.Sleep(time.Second * 1)
		minutesPassed++
		t := 5 - minutesPassed
		sendNotification("Pomodoro timer\nBreak left: " + strconv.FormatInt(t, 10))
	}
	sendNotification("Pomodoro timer\nBreak ended.\nStarting new pomodoro cycle!")
}

func handler() {
	flag := os.Args[1]
	switch flag {
	case "start":
		var cycles int64 = 0
		for {
			startPomodoro(cycles)
			cycles++
		}
	default:
		fmt.Println("Invalid flag!")
		break
	}
}

func main() {
	handler()
}
