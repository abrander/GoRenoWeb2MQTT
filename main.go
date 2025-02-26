package main

import (
	"os"
	"time"

	"github.com/anderskvist/GoRenoWeb2MQTT/mqtt"
	"github.com/anderskvist/GoRenoWeb2MQTT/renoweb"

	"github.com/anderskvist/DVIEnergiSmartControl/log"
	ini "gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load(os.Args[1])

	if err != nil {
		log.Criticalf("Fail to read file: %v", err)
		os.Exit(1)
	}

	poll := cfg.Section("main").Key("poll").MustInt(60)
	log.Infof("Polltime is %d seconds.\n", poll)
	for t := range time.NewTicker(time.Duration(poll) * time.Second).C {
		log.Notice("Tick")
		if t == t {
		}
		log.Info("Getting data from RenoWeb")
		addressID := renoweb.GetRenoWebAddressID(cfg.Section("renoweb").Key("address").String())
		pickupPlans := renoweb.GetRenoWebPickupPlan(addressID)
		log.Info("Done getting data from RenoWeb")

		log.Info("Sending data to MQTT")
		mqtt.SendToMQTT(cfg, pickupPlans)
		log.Info("Done sending to MQTT")
	}
}
