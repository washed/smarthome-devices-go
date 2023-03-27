package smarthome_devices

import (
	"encoding/json"
	"fmt"
	"testing"
)

const ShellyTRVInfoJSON = `
{
    "thermostats": [
        {
            "pos": -100.0,
            "target_t": {
                "enabled": true,
                "value": 31.0,
                "units": "C"
            },
            "tmp": {
                "value": 17.4,
                "units": "C",
                "is_valid": true
            },
            "schedule": false,
            "schedule_profile": 2,
            "boost_minutes": 0
        }
    ],
    "calibrated": false,
    "bat": {
        "value": 0,
        "voltage": 3.127
    },
    "charger": true,
    "ps_mode": 0,
    "dbg_flags": 0
}
`

const ShellyTRVInfoJSON_Broken = "{\"wifi_sta\":{\"connected\":true,\"ssid\":\"wpd.wlan-2.4GHz\",\"ip\":\"192.168.178.82\",\"rssi\":-51},\"cloud\":{\"enabled\":false,\"connected\":false},\"mqtt\":{\"connected\":true},\"time\":\"14:41\",\"unixtime\":1676900460,\"serial\":0,\"has_update\":false,\"mac\":\"60A423D92566\",\"cfg_changed_cnt\":0,\"actions_stats\":{\"skipped\":0},\"thermostats\":[{\"pos\":0.0,\"target_t\":{\"enabled\":true,\"value\":23.0,\"value_op\":8.0,\"units\":\"C\"},\"tmp\":{\"value\":22.1,\"units\":\"C\",\"is_valid\":true},\"schedule\":true,\"schedule_profile\":1,\"boost_minutes\":0,\"window_open\":false}],\"calibrated\":true,\"bat\":{\"value\":71,\"voltage\":3.665},\"charger\":false,\"update\":{\"status\":\"unknown\",\"has_update\":false,\"new_version\":\"20220811-152343/v2.1.8@5afc928c\",\"old_version\":\"20220811-152343/v2.1.8@5afc928c\",\"beta_version\":null},\"ram_total\":97280,\"ram_free\":22568,\"fs_size\":65536,\"fs_free\":59440,\"uptime\":43083,\"fw_info\":{\"device\":\"shellytrv-60A423D92566\",\"fw\":\"20220811-152343/v2.1.8@5afc928c\"},\"ps_mode\":0,\"dbg_flags\":0}\x00zautodetect\":true,\"tz_utc_offset\":3600,\"tz_dst\":false,\"tz_dst_auto\":true,\"time\":\"14:40\",\"child_lock\":false,\"clog_prevention\":false,\"display\":{\"brightness\":7,\"flipped\":false},\"hwinfo\":{\"hw_revision\":\"dev-prototype\",\"batch_id\":0},\"sleep_mode\":{\"period\":60,\"unit\":\"m\"},\"thermostats\":[{\"target_t\":{\"enabled\":true,\"value\":23.0,\"value_op\":8.0,\"units\":\"C\",\"accelerated_heating\":true},\"schedule\":true,\"schedule_profile\":1,\"schedule_profile_names\":[\"Livingroom\",\"Livingroom 1\",\"Bedroom\",\"Bedroom 1\",\"Holiday\"],\"schedule_rules\":[\"0600-0123456-23\",\"2300-0123456-18\"],\"temperature_offset\":0.0,\"ext_t\":{\"enabled\":false, \"floor_heating\": false},\"t_auto\":{\"enabled\":true},\"boost_minutes\":30,\"valve_min_percent\":0.00,\"force_close\":false,\"calibration_correction\":true,\"extra_pressure\":false,\"open_window_report\":true}] }"

func TestShellyTRVDatumFromJSON(t *testing.T) {
	sd := ShellyTRVInfo{}
	err := json.Unmarshal([]byte(ShellyTRVInfoJSON), &sd)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("sd: %+v", sd)
}
