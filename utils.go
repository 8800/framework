package utils

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Debug() {
	fmt.Println("这是库打印")
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("这是库-log打印")
}
