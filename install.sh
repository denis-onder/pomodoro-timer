#!/bin/bash

# Get dependencies
go get -u github.com/faiface/beep
go get -u	github.com/pkg/errors
go get -u github.com/hajimehoshi/go-mp3
go get -u github.com/hajimehoshi/oto

# Build and move the executable to /usr/local/bin
rm pomodoro-timer # Remove any old executables laying around
go build .
sudo rm /usr/local/bin/pomodoro-timer # Remove any old executables
sudo mv ./pomodoro-timer /usr/local/bin

# Copy sound file to where the executable is
sudo rm /usr/local/bin/ding.mp3 # Remove any old sound files
sudo cp ./ding.mp3 /usr/local/bin

# Notify user that it's all done
echo "Moved the binary to /usr/local/bin"
echo "You can now use pomodoro-timer from anywhere in your terminal!"