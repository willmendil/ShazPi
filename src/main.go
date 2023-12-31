package main

import (

	// "shazammini/src/microphone"

	"shazammini/src/commands"
	"shazammini/src/io"
	"shazammini/src/structs"
	"time"

	"github.com/d2r2/go-logger"
	"gobot.io/x/gobot"
)

func main() {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	io.New()
	defer io.Kill()

	master := gobot.NewMaster()

	commCahnnels := structs.CommChannels{
		PlayChannel:     make(chan bool),
		RecordChannel:   make(chan time.Duration),
		FetchAPI:        make(chan bool),
		DisplayResult:   make(chan structs.Track),
		DisplayRecord:   make(chan bool),
		DisplayThinking: make(chan bool),
	}

	_ = commCahnnels

	// dis := display.Screen(&commCahnnels)
	// mic := microphone.Microphone(&commCahnnels)
	com := commands.Commands(&commCahnnels)
	// api := api.Api(&commCahnnels)

	// master.AddRobot(dis)
	// master.AddRobot(api)
	master.AddRobot(com)
	// master.AddRobot(mic)

	master.Start()
}
