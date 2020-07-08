package generator

import (
	btc_client "btc-faucet.drozd.by/modules/btc-client"
	"fmt"
	"log"
	"time"
)

var isGeneratorEnabled = false
var generatorInterval = 1 * time.Minute
var stopChannel = make(chan bool)
var addressGenerator = ""

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

func GetGeneratorInterval() time.Duration {
	return generatorInterval
}

func doWork() {
	isGeneratorEnabled = true
	go func() {
		for {
			select {
			case <-time.After(generatorInterval):
				GenerateBlock()
				fmt.Println("tick block generate")
			case <-stopChannel:
				fmt.Println("Stopping generator")
				isGeneratorEnabled = false
				return
			}
		}
	}()
}

func GenerateBlock() {
	addressGenerator = GetRefundAddress()
	err := btc_client.GenerateToAddress(1, addressGenerator)
	if err != nil {
		log.Fatal(err)
	}
}

func GetRefundAddress() string {
	if addressGenerator == "" {
		address, err := btc_client.GetAddressByLabelOrNew("mining")
		if err != nil {
			log.Fatal(err)
		}
		addressGenerator = address
	}

	return addressGenerator
}