package rabbitmq

import (
	"log"
)

// helper

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
