package smarthome_devices

import (
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
)

type ShellySwitch25 struct {
	MqttDevice
}

func NewShellySwitch25(deviceId string, mqttOpts *MQTT.ClientOptions) ShellySwitch25 {
	client := MQTT.NewClient(mqttOpts)
	s := ShellySwitch25{MqttDevice{DeviceId: deviceId, mqttClient: client, mqttOpts: mqttOpts}}
	log.Debug().Str("DeviceName", s.DeviceName()).Msg("New ShellySwitch25")
	return s
}

func (s ShellySwitch25) Connect() {
	if token := s.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Error().
			Str("DeviceName", s.DeviceName()).
			Err(token.Error()).
			Msg("Error connecting to MQTT!")
	}
	log.Info().Str("DeviceName", s.DeviceName()).Msg("connected")
}

func (s ShellySwitch25) Close() {
	s.mqttClient.Disconnect(disconnectQiesceTimeMs)
	log.Info().Str("DeviceName", s.DeviceName()).Msg("disconnected")
}

func (s ShellySwitch25) DeviceName() string {
	return fmt.Sprintf("shellyswitch25-%s", s.DeviceId)
}

func (s ShellySwitch25) baseTopic() string {
	return fmt.Sprintf("shellies/%s", s.DeviceName())
}

func (s ShellySwitch25) baseCommandTopic() string {
	return s.baseTopic() + "/roller/0/command"
}

/*
func (s ShellySwitch25) SubscribeRelayState(onHandler func(), offHandler func()) {
	topic := s.baseTopic() + "/relay/0"
	relayStateCallback := func(relayState string) {
		if relayState == "on" {
			onHandler()
		} else if relayState == "off" {
			offHandler()
		} else {
			log.Error().Str("relayState", relayState).Msg("received unknown relayState value")
		}
	}

	SubscribeStringHelper(s.mqttClient, topic, relayStateCallback)
}

func (s ShellySwitch25) SubscribePower(powerHandler func(float32)) {
	topic := s.baseTopic() + "/relay/0/power"
	powerCallback := func(powerStr string) {
		power, err := strconv.ParseFloat(powerStr, 32)
		if err != nil {
			log.Error().Str("powerStr", powerStr).Msg("error parsing powerStr as float32")
		}
		powerHandler(float32(power))
	}

	SubscribeStringHelper(s.mqttClient, topic, powerCallback)
}
*/

func (s ShellySwitch25) rollerCommand(command string) {
	log.Info().
		Str("DeviceName", s.DeviceName()).
		Str("command", command).
		Msg("sending roller command")
	topic := s.baseCommandTopic()

	token := s.mqttClient.Publish(topic, byte(qos), false, command)
	token.Wait()
}

func (s ShellySwitch25) RollerOpen() {
	s.rollerCommand("open")
}

func (s ShellySwitch25) RollerCalibrate() {
	s.rollerCommand("rc")
}

func (s ShellySwitch25) RollerClose() {
	s.rollerCommand("close")
}

func (s ShellySwitch25) RollerStop() {
	s.rollerCommand("stop")
}

func (s ShellySwitch25) RollerSetPosition(pos int) {
	if pos < 0 || pos > 100 {
		log.Error().Int("pos", pos).Msg("position must be between 0 and 100")
		return
	}

	log.Info().
		Str("DeviceName", s.DeviceName()).
		Int("pos", pos).
		Msg("sending roller positino command")
	topic := s.baseCommandTopic() + "/pos"
	token := s.mqttClient.Publish(topic, byte(qos), false, fmt.Sprintf("%d", pos))
	token.Wait()
}
