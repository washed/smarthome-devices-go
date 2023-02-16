package smarthome_devices

/*
Aqara (Xiami) WSDCGQ11LM temperature, humidity, and pressure sensor
*/

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type AqaraTHP struct {
	MqttDevice
}

/*
ZigBee2MQTT payload example:

{
    "battery": 30,
    "humidity": 42.97,
    "linkquality": 215,
    "power_outage_count": 965,
    "pressure": 966.2,
    "temperature": 21.75,
    "voltage": 2895
}
*/

type AqaraTHPStatus struct {
	Battery          int16   `json:"battery"`
	Humidity         float32 `json:"humidity"`
	Linkquality      uint8   `json:"linkquality"`
	PowerOutageCount int     `json:"power_outage_count"`
	Pressure         float32 `json:"pressure"`
	Temperature      float32 `json:"temperature"`
	Voltage          int     `json:"voltage"`
}

func NewAqaraTHP(deviceId string, mqttOpts *MQTT.ClientOptions) AqaraTHP {
	client := MQTT.NewClient(mqttOpts)
	d := AqaraTHP{MqttDevice{DeviceId: deviceId, mqttClient: client, mqttOpts: mqttOpts}}
	log.Debug().Str("DeviceName", d.DeviceName()).Msg("New AqaraTHP")

	return d
}

func (s AqaraTHP) Connect() {
	if token := s.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Error().
			Str("DeviceName", s.DeviceName()).
			Err(token.Error()).
			Msg("Error connecting to MQTT!")
	}
	log.Info().Str("DeviceName", s.DeviceName()).Msg("connected")
}

func (s AqaraTHP) Close() {
	s.mqttClient.Disconnect(disconnectQiesceTimeMs)
	log.Info().Str("DeviceName", s.DeviceName()).Msg("disconnected")
}

func (s AqaraTHP) DeviceName() string {
	return fmt.Sprintf("aqara-thp/%s", s.DeviceId)
}

func (s AqaraTHP) baseTopic() string {
	return fmt.Sprintf("zigbee2mqtt/%s", s.DeviceName())
}

type AqaraTHPStatusCallback = func(status AqaraTHPStatus)

func (s AqaraTHP) SubscribeStatus(statusCallback AqaraTHPStatusCallback) {
	SubscribeJSONHelper(s.mqttClient, s.baseTopic(), statusCallback)
}
