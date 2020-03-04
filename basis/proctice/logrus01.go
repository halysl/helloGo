package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func logrusDemo() {
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{
		"where": "here",
		"when": "now",
	}).Info("simple log")

	log.SetLevel(log.ErrorLevel)
	log.Debug("Debug")
	log.Info("Info ")
	log.Error("Error could appear")

	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.JSONFormatter{})
	log.WithField("animal", "dog").Info("animal include dog")
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.WithField("animal", "dog").Info("animal include dog")
}
