package lib

import (
	"github.com/nsqio/go-nsq"
	"github.com/prometheus/common/log"
	"os"
)

var isInited = false
var producer *nsq.Producer

func GetNsqProducer() *nsq.Producer {
	if !isInited {
		config := nsq.NewConfig()
		var err error
		producer, err = nsq.NewProducer(os.Getenv("NSQ_URI"), config)
		if err != nil {
			log.Fatal("Can not init nsq producer", err)
		}
		isInited = true
	}

	return producer
}