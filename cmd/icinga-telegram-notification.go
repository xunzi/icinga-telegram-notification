package main

import (
	"flag"
	"log"
	"os"
	telegramnnotifier "telegramnotifier"
)

var botToken = flag.String("botToken", "", "Bot Token")
var objectType = flag.String("objectType", "service", "Object type (host or service)")
var chatID = flag.String("chatID", "", "Chat Id")
var notificationType = flag.String("notificationType", "", "Notification type")
var hostName = flag.String("hostName", "", "host name")
var state = flag.String("state", "", "State (up, down)")
var serviceName = flag.String("serviceName", "", "Service Name")
var outPut = flag.String("outPut", "", "Check command output")
var timeStamp = flag.String("timeStamp", "", "Event time stamp")
var debug = flag.Bool("debug", false, "Debug/verbose mode")
var notificationMessage string
var logFile = flag.String("logFile", "", "Log file")

func main() {
	flag.Parse()
	telegramnnotifier.Debug = *debug
	if *logFile != "" {
		logfh, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening logfile %s: %v", *logFile, err)
		}
		log.SetOutput(logfh)
	}
	notificationMessage = telegramnnotifier.GenerateNotification(*objectType, *notificationType, *hostName,
		*serviceName, *state, *outPut, *timeStamp)
	telegramnnotifier.SendNotification(*botToken, *chatID, notificationMessage)
}
