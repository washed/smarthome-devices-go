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

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano},
	)

	mqttOpts := MQTT.NewClientOptions()
	mqttOpts.AddBroker(broker)
	mqttOpts.SetUsername(user)
	mqttOpts.SetPassword(password)

	switch25 := shd.NewShellySwitch25("40F520051602", mqttOpts)
	switch25.Connect()
	defer switch25.Close()

	switch25.RollerSetPosition(6)
}
