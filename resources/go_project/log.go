package go_project

import (
	log "github.com/sirupsen/logrus"
)

func Log() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
