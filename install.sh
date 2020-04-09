#!/bin/bash

# Get dependencies
go get -u github.com/faiface/beep
go get -u	github.com/pkg/errors
go get -u github.com/hajimehoshi/go-mp3
go get -u github.com/hajimehoshi/oto

# Build and move the executable to /usr/local/bin
rm pomodoro-timer # Remove any old executables laying around
go build .
sudo mv ./pomodoro-timer /usr/local/bin

# Notify user that it's all done
echo "Moved the binary to /usr/local/bin"
echo "You can now use pomodoro-timer from anywhere in your terminal!"