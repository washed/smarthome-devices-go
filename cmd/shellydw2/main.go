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

func infoCallback(info shd.ShellyDW2Info) {
	log.Info().
		Interface("info", info).
		Msg("Received ShellyDW2Info")
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

	dw2 := shd.NewShellyDW2("C92B94", mqttOpts)
	dw2.Connect()
	defer dw2.Close()

	dw2.SubscribeInfo(infoCallback)

	for {
		time.Sleep(time.Second * 10)
	}
}
