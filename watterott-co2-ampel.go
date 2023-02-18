package smarthome_devices

/*
Watterott CO2 Ampel (washed FW)
*/

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type WatterottCO2Ampel struct {
	MqttDevice
}

/*
status payload example:

{
	"co2": 1182,
	"temp": 10.37269592,
	"hum": 83.64657593,
	"lux": 70
}
*/

type WatterottCO2AmpelStatus struct {
	CO2         int32   `json:"co2"`
	Temperature float32 `json:"temp"`
	Humidity    float32 `json:"hum"`
	Lux         int32   `json:"lux"`
}

func NewWatterottCO2Ampel(deviceId string, mqttOpts *MQTT.ClientOptions) WatterottCO2Ampel {
	client := MQTT.NewClient(mqttOpts)
	d := WatterottCO2Ampel{MqttDevice{DeviceId: deviceId, mqttClient: client, mqttOpts: mqttOpts}}
	log.Debug().Str("DeviceName", d.DeviceName()).Msg("New WatterottCO2Ampel")

	return d
}

func (s WatterottCO2Ampel) Connect() {
	if token := s.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Error().
			Str("DeviceName", s.DeviceName()).
			Err(token.Error()).
			Msg("Error connecting to MQTT!")
	}
	log.Info().Str("DeviceName", s.DeviceName()).Msg("connected")
}

func (s WatterottCO2Ampel) Close() {
	s.mqttClient.Disconnect(disconnectQiesceTimeMs)
	log.Info().Str("DeviceName", s.DeviceName()).Msg("disconnected")
}

func (s WatterottCO2Ampel) DeviceName() string {
	return fmt.Sprintf("sensors/%s", s.DeviceId)
}

func (s WatterottCO2Ampel) baseTopic() string {
	return fmt.Sprintf("sensors/%s", s.DeviceId)
}

type WatterottCO2AmpelStatusCallback = func(status WatterottCO2AmpelStatus)

func (s WatterottCO2Ampel) SubscribeStatus(statusCallback WatterottCO2AmpelStatusCallback) {
	SubscribeJSONHelper(s.mqttClient, s.baseTopic(), statusCallback)
}
