package smarthome_devices

import (
	"testing"
)

type Message struct {
	payload []byte
}

func (m Message) Payload() []byte {
	return m.payload
}

func (m Message) Duplicate() bool   { return false }
func (m Message) Qos() byte         { return 0 }
func (m Message) Retained() bool    { return false }
func (m Message) Topic() string     { return "" }
func (m Message) MessageID() uint16 { return 0 }
func (m Message) Ack()              {}

func TestCheckedJSONUnmarshal(t *testing.T) {
	output := map[string]interface{}{}
	m := Message{payload: []byte(ShellyTRVInfoJSON_Broken)}

	err := checkedJSONUnmarshal(m, &output)
	if err != nil {
		t.Errorf("%s", err)
	}
}
