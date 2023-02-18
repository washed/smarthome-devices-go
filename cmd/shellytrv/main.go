package main

import (
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	shd "github.com/washed/smarthome-devices-go"
)

var (
	broker   = os.Getenv("MQTT_BROKER_URL")
	user     = os.Getenv("MQTT_BROKER_USERNAME")
	password = os.Getenv("MQTT_BROKER_PASSWORD")
)

func infoCallback(info shd.ShellyTRVInfo) {
	log.Info().
		Interface("info", info).
		Msg("Received ShellyTRVInfo")
}

func statusCallback(status shd.ShellyTRVStatus) {
	log.Info().
		Interface("status", status).
		Msg("Received ShellyTRVStatus")
}

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano},
	)

	mqttOpts := MQTT.NewClientOptions()
	mqttOpts.AddBroker(broker)
	mqttOpts.SetUsername(user)
	mqttOpts.SetPassword(password)

	trv := shd.NewShellyTRV("60A423DAE8DE", mqttOpts)
	trv.Connect()
	defer trv.Close()

	trv.SubscribeAll()
	trv.SubscribeInfo(infoCallback)
	trv.SubscribeStatus(statusCallback)

	for {
		time.Sleep(time.Second * 10)
	}
}
