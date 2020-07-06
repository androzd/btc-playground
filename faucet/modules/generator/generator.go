package generator

import (
	btc_client "btc-faucet.drozd.by/modules/btc-client"
	"fmt"
	"log"
	"time"
)

var isGeneratorEnabled = false
var generatorInterval = 1 * time.Second
var stopChannel = make(chan bool)

func StartGeneratorRoutine() {
	if !isGeneratorEnabled {
		doWork()
	}
}

func StopGeneratorRoutine() {
	if isGeneratorEnabled {
		stopChannel <- true
	}
}

func StatusGeneratorRoutine() bool {
	return isGeneratorEnabled
}

func SetGeneratorInterval(duration time.Duration) {
	generatorInterval = duration
}

func doWork() {
	isGeneratorEnabled = true
	go func() {
		address, err := btc_client.GetAddressByLabelOrNew("mining")
		if err != nil {
			log.Fatal(err)
		}
		for {
			select {
			case <-time.After(generatorInterval):
				err = btc_client.GenerateToAddress(1, address)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("tick block generate")
			case <-stopChannel:
				fmt.Println("Stopping generator")
				isGeneratorEnabled = false
				return
			}
		}
	}()
}