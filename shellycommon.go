package smarthome_devices

type ShellyInfoBat struct {
	Value   int     `json:"value"`
	Voltage float32 `json:"voltage"`
}

type ShellyInfoTmp struct {
	Value   float32 `json:"value"`
	Units   string  `json:"units"`
	IsValid bool    `json:"is_valid"`
}

type ShellyInfoLux struct {
	Value        float32 `json:"value"`
	Illumination string  `json:"illumination"`
	IsValid      bool    `json:"is_valid"`
}
