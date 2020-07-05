package services

import (
	"fmt"
	"time"
)

var isGeneratorEnabled = false
var generatorInterval = 1 * time.Second
var stopChannel = make(chan bool)

func StartGeneratorRoutine() {
	doWork()
}

func StopGeneratorRoutine() {
	stopChannel <- true
}

func SetGeneratorInterval(duration time.Duration) {
	generatorInterval = duration
}

func doWork() {
	if isGeneratorEnabled {
		fmt.Println( "Generator already enabled")
		return
	}

	isGeneratorEnabled = true
	go func() {
		for {
			select {
			case <-time.After(generatorInterval):
				fmt.Println("tick block generate")
			case <- stopChannel:
				fmt.Println("Stopping generator")
				return
			}
		}
	}()
	isGeneratorEnabled = false
}