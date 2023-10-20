package models

import "fmt"

type Attributes struct {
	CommandTemplate   string   `json:"command_template"`
	DeviceClass       string   `json:"device_class"`
	EntityType        string   `json:"entity_type"`
	Max               float64  `json:"max"`
	Min               float64  `json:"min"`
	Name              string   `json:"name"`
	Options           []string `json:"options"`
	StateClass        string   `json:"state_class"`
	Step              float64  `json:"step"`
	UnitOfMeasurement string   `json:"unit_of_measurement"`
	ValueTemplate     string   `json:"value_template"`
}

type ParamsMap map[string]Attributes

func (m ParamsMap) AddEnabledParams(from map[string]string, isReadOnly bool) {
	m.addLastUpdate(from, true, isReadOnly, !isReadOnly)
	m.addGlobals(from, true, isReadOnly, !isReadOnly)
	m.addDomesticHotWater(from, true, isReadOnly, !isReadOnly)
	m.addAnalogInput(from, true, isReadOnly, !isReadOnly)
	m.addDigitalInput(from, true, isReadOnly, !isReadOnly)
	m.addDigitalAlarm(from, true, isReadOnly, !isReadOnly)
	m.addZones(from, true, isReadOnly, !isReadOnly)
	m.addCircuits(from, true, isReadOnly, !isReadOnly)
	m.addSources(from, true, isReadOnly, !isReadOnly)
	m.addDehumidifier(from, true, isReadOnly, !isReadOnly)
	m.addEnergymeters(from, true, isReadOnly, !isReadOnly)
}

func (m ParamsMap) AddDisabledParams(from map[string]string, isReadOnly bool) {
	m.addGlobals(from, false, !isReadOnly, isReadOnly)
	m.addDomesticHotWater(from, false, !isReadOnly, isReadOnly)
	m.addAnalogInput(from, false, !isReadOnly, isReadOnly)
	m.addDigitalInput(from, false, !isReadOnly, isReadOnly)
	m.addDigitalAlarm(from, false, !isReadOnly, isReadOnly)
	m.addZones(from, false, !isReadOnly, isReadOnly)
	m.addCircuits(from, false, !isReadOnly, isReadOnly)
	m.addSources(from, false, !isReadOnly, isReadOnly)
	m.addDehumidifier(from, false, !isReadOnly, isReadOnly)
	m.addEnergymeters(from, false, !isReadOnly, isReadOnly)
}

func (m ParamsMap) addLastUpdate(from map[string]string, static, read, write bool) {
	if static {
		m["LAST_UPDATE"] = Attributes{
			DeviceClass: "timestamp",
			EntityType:  "sensor",
			Name:        "Last update",
		}
	}
}

func (m ParamsMap) addGlobals(from map[string]string, static, read, write bool) {
	if static {
		m["GLOBAL_ENABLE"] = Attributes{
			Name:          "Global state",
			EntityType:    "binary_sensor",
			ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
		}
		m["GLOBAL_T_EXT"] = Attributes{
			Name:              "Global external temperature",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
		m["GLOBAL_SEASON"] = Attributes{
			Name:          "Global season",
			EntityType:    "sensor",
			DeviceClass:   "enum",
			ValueTemplate: "{% if value == \"0\" %}winter{% elif value == \"1\" %}summer{% else %}{{ value }}{% endif %}",
		}
		m["GLOBAL_DEICING"] = Attributes{
			Name:          "Global de-ice state",
			EntityType:    "binary_sensor",
			ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
		}
		m["GLOBAL_EXPECTED_DEWP"] = Attributes{
			Name:              "Global dewpoint",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
	}
	if read {
		m["GLOBAL_ZONE_T_HYST"] = Attributes{
			Name:              "Global zone temperature hysteresis",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
		m["GLOBAL_ZONE_RH_HYST"] = Attributes{
			Name:              "Global zone humidity hysteresis",
			EntityType:        "sensor",
			DeviceClass:       "humidity",
			UnitOfMeasurement: "%",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
		m["GLOBAL_ZONE_DEICE_TRESH"] = Attributes{
			Name:              "Global zone de-ice threshold",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
	}
	if write {
		m["GLOBAL_ZONE_T_HYST"] = Attributes{
			Name:              "Global zone temperature hysteresis",
			EntityType:        "number",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			Max:               1,
			Min:               0.1,
			Step:              0.1,
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
			CommandTemplate:   "{{ (value * 10) | int }}",
		}
		m["GLOBAL_ZONE_RH_HYST"] = Attributes{
			Name:              "Global zone humidity hysteresis",
			EntityType:        "number",
			DeviceClass:       "humidity",
			UnitOfMeasurement: "%",
			Max:               5,
			Min:               1,
			Step:              0.1,
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
			CommandTemplate:   "{{ (value * 10) | int }}",
		}
		m["GLOBAL_ZONE_DEICE_TRESH"] = Attributes{
			Name:              "Global zone de-ice threshold",
			EntityType:        "number",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			Max:               10,
			Min:               6,
			Step:              0.1,
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
			CommandTemplate:   "{{ (value * 10) | int }}",
		}
	}
}

func (m ParamsMap) addDomesticHotWater(from map[string]string, static, read, write bool) {
	if static {
		m["ACS_MAIN_OUTPUT"] = Attributes{
			Name:          "DHW state",
			EntityType:    "binary_sensor",
			ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
		}
		m["GLOBAL_ACS_ENABLE"] = Attributes{
			Name:          "DHW enabled",
			EntityType:    "binary_sensor",
			ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
		}
		m["GLOBAL_T_ACS"] = Attributes{
			Name:              "DHW temperature",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
		m["GLOBAL_SET_ACS"] = Attributes{
			Name:              "DHW temperature setpoint",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
	}
	if read {
		m["ACS_SET_ECONOMY"] = Attributes{
			Name:              "DHW economy setpoint",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
		m["ACS_SET_COMFORT"] = Attributes{
			Name:              "DHW comfort setpoint",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
		m["ACS_SET_HYST"] = Attributes{
			Name:              "DHW setpoint hysteresis",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
		m["ACS_SET_DELTA"] = Attributes{
			Name:              "DHW second stage deviation",
			EntityType:        "sensor",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
		}
	}
	if write {
		m["ACS_SET_ECONOMY"] = Attributes{
			Name:              "DHW economy setpoint",
			EntityType:        "number",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			Max:               60,
			Min:               30,
			Step:              0.1,
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
			CommandTemplate:   "{{ (value * 10) | int }}",
		}
		m["ACS_SET_COMFORT"] = Attributes{
			Name:              "DHW comfort setpoint",
			EntityType:        "number",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			Max:               60,
			Min:               30,
			Step:              0.1,
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
			CommandTemplate:   "{{ (value * 10) | int }}",
		}
		m["ACS_SET_HYST"] = Attributes{
			Name:              "DHW setpoint hysteresis",
			EntityType:        "number",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			Max:               10,
			Min:               0,
			Step:              0.1,
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
			CommandTemplate:   "{{ (value * 10) | int }}",
		}
		m["ACS_SET_DELTA"] = Attributes{
			Name:              "DHW second stage deviation",
			EntityType:        "number",
			DeviceClass:       "temperature",
			UnitOfMeasurement: "°C",
			Max:               10,
			Min:               0,
			Step:              0.1,
			StateClass:        "measurement",
			ValueTemplate:     "{{ value | int / 10 }}",
			CommandTemplate:   "{{ (value * 10) | int }}",
		}
	}
}

func (m ParamsMap) addAnalogInput(from map[string]string, static, read, write bool) {
	if static {
		for i := 1; i <= 8; i++ {
			if from["FAIN"+fmt.Sprint(i)+"_TEMP"] != "32769" {
				m["FAIN"+fmt.Sprint(i)+"_TEMP"] = Attributes{
					Name:              "Analog input " + fmt.Sprint(i),
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					StateClass:        "measurement",
					UnitOfMeasurement: "°C",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "",
				}
			}
		}
	}
}

func (m ParamsMap) addDigitalInput(from map[string]string, static, read, write bool) {
	if static {
		for i := 1; i <= 8; i++ {
			m["FDIN"+fmt.Sprint(i)+"_STATUS"] = Attributes{
				Name:          "Digital input " + fmt.Sprint(i),
				EntityType:    "binary_sensor",
				ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
			}
		}
	}
}

func (m ParamsMap) addDigitalAlarm(from map[string]string, static, read, write bool) {
	if static {
		for i := 1; i <= 5; i++ {
			m["FALDIN"+fmt.Sprint(i)+"_STATUS"] = Attributes{
				Name:          "Alarm " + fmt.Sprint(i),
				EntityType:    "binary_sensor",
				ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
			}
		}
	}
}

func (m ParamsMap) addZones(from map[string]string, static, read, write bool) {
	for i := 1; i <= 32; i++ {
		if from["Z"+fmt.Sprint(i)+"_TEMP"] != "32769" {
			if static {
				m["Z"+fmt.Sprint(i)+"_OUTPUT"] = Attributes{
					Name:          "Zone " + fmt.Sprint(i) + " state",
					EntityType:    "binary_sensor",
					ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
				}
				m["Z"+fmt.Sprint(i)+"_TEMP"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " temperature",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["Z"+fmt.Sprint(i)+"_ZONE_MODE"] = Attributes{
					Name:          "Zone " + fmt.Sprint(i) + " mode",
					EntityType:    "sensor",
					DeviceClass:   "enum",
					ValueTemplate: "{% if value == \"0\" %}off{% elif value == \"2\" %}economy{% elif value == \"3\" %}comfort{% elif value == \"4\" %}forced off{% elif value == \"6\" %}forced economy{% elif value == \"23\" %}forced comfort{% else %}{{ value }}{% endif %}",
				}
				m["Z"+fmt.Sprint(i)+"_ZONE_SET"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " setpoint",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
			}
			if read {
				m["Z"+fmt.Sprint(i)+"_FORCING"] = Attributes{
					Name:          "Zone " + fmt.Sprint(i) + " preset",
					EntityType:    "sensor",
					DeviceClass:   "enum",
					ValueTemplate: "{% if value == \"0\" %}automatic{% elif value == \"1\" %}forced off{% elif value == \"2\" %}forced economy{% elif value == \"3\" %}forced comfort{% else %}{{ value }}{% endif %}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_CW"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " C.W. setpoint",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_EW"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " E.W. setpoint",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_CS"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " C.S. setpoint",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_ES"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " E.S. setpoint",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
			}
			if write {
				m["Z"+fmt.Sprint(i)+"_FORCING"] = Attributes{
					Name:            "Zone " + fmt.Sprint(i) + " preset",
					Options:         []string{"automatic", "forced off", "forced economy", "forced comfort"},
					EntityType:      "select",
					ValueTemplate:   "{% if value == \"1\" %}forced off{% elif value == \"2\" %}forced economy{% elif value == \"3\" %}forced comfort{% else %}automatic{% endif %}",
					CommandTemplate: "{% if value == \"forced off\" %}1{% elif value == \"forced economy\" %}2{% elif value == \"forced comfort\" %}3{% else %}0{% endif %}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_CW"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " C.W. setpoint",
					EntityType:        "number",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					Max:               30,
					Min:               15,
					Step:              0.1,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_EW"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " E.W. setpoint",
					EntityType:        "number",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					Max:               30,
					Min:               15,
					Step:              0.1,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_CS"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " C.S. setpoint",
					EntityType:        "number",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					Max:               30,
					Min:               15,
					Step:              0.1,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["Z"+fmt.Sprint(i)+"_SET_ES"] = Attributes{
					Name:              "Zone " + fmt.Sprint(i) + " E.S. setpoint",
					EntityType:        "number",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					Max:               30,
					Min:               15,
					Step:              0.1,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
			}
			if from["Z"+fmt.Sprint(i)+"_RH"] != "32769" {
				if static {
					m["Z"+fmt.Sprint(i)+"_RH"] = Attributes{
						Name:              "Zone " + fmt.Sprint(i) + " humidity",
						EntityType:        "sensor",
						DeviceClass:       "humidity",
						UnitOfMeasurement: "%",
						StateClass:        "measurement",
						ValueTemplate:     "{{ value | int / 10 }}",
						CommandTemplate:   "{{ (value * 10) | int }}",
					}
				}
				if read {
					m["Z"+fmt.Sprint(i)+"_SET_RH"] = Attributes{
						Name:              "Zone " + fmt.Sprint(i) + " humidity setpoint",
						EntityType:        "sensor",
						DeviceClass:       "humidity",
						UnitOfMeasurement: "%",
						StateClass:        "measurement",
						ValueTemplate:     "{{ value | int / 10 }}",
						CommandTemplate:   "{{ (value * 10) | int }}",
					}
				}
				if write {
					m["Z"+fmt.Sprint(i)+"_SET_RH"] = Attributes{
						Name:              "Zone " + fmt.Sprint(i) + " humidity setpoint",
						EntityType:        "number",
						DeviceClass:       "humidity",
						UnitOfMeasurement: "%",
						Max:               70,
						Min:               40,
						Step:              0.1,
						StateClass:        "measurement",
						ValueTemplate:     "{{ value | int / 10 }}",
						CommandTemplate:   "{{ (value * 10) | int }}",
					}
				}
			}
		}
	}
}

func (m ParamsMap) addCircuits(from map[string]string, static, read, write bool) {
	for i := 1; i <= 8; i++ {
		if from["C"+fmt.Sprint(i)+"_TEMP"] != "32769" {
			if static {
				m["C"+fmt.Sprint(i)+"_TEMP"] = Attributes{
					Name:              "Circuit " + fmt.Sprint(i) + " temperature",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
				m["C"+fmt.Sprint(i)+"_SET"] = Attributes{
					Name:              "Circuit " + fmt.Sprint(i) + " temperature setpoint",
					EntityType:        "sensor",
					DeviceClass:       "temperature",
					UnitOfMeasurement: "°C",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "{{ (value * 10) | int }}",
				}
			}
		}
	}
}

func (m ParamsMap) addSources(from map[string]string, static, read, write bool) {
	for i := 1; i <= 3; i++ {
		if from["S"+fmt.Sprint(i)+"_DESCR"] != "0" {
			if static {
				m["S"+fmt.Sprint(i)+"_ENABLED"] = Attributes{
					Name:          "Source " + fmt.Sprint(i) + " enabled",
					EntityType:    "binary_sensor",
					ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
				}
				m["S"+fmt.Sprint(i)+"_OUTPUT"] = Attributes{
					Name:          "Source " + fmt.Sprint(i) + " state",
					EntityType:    "binary_sensor",
					ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
				}
				m["S"+fmt.Sprint(i)+"_AUXOUTPUT"] = Attributes{
					Name:          "Source " + fmt.Sprint(i) + " auxiliary state",
					EntityType:    "binary_sensor",
					ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
				}
			}
		}
	}
}

func (m ParamsMap) addDehumidifier(from map[string]string, static, read, write bool) {
	for i := 1; i <= 8; i++ {
		if from["D"+fmt.Sprint(i)+"_SPEED_LOW"] != "0" && from["D"+fmt.Sprint(i)+"_SPEED_ECONOMY"] != "0" {
			if static {
				m["D"+fmt.Sprint(i)+"_OUTPUT_RENEW"] = Attributes{
					Name:          "Fan " + fmt.Sprint(i) + " renew",
					EntityType:    "binary_sensor",
					ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
				}
				m["D"+fmt.Sprint(i)+"_OUTPUT_DEUM"] = Attributes{
					Name:          "Fan " + fmt.Sprint(i) + " dehumidify",
					EntityType:    "binary_sensor",
					ValueTemplate: "{% if value == \"1\" %}on{% else %}off{% endif %}",
				}

			}
			if read {
				m["D"+fmt.Sprint(i)+"_SPEED_LOW"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " low flow rate",
					EntityType:        "sensor",
					UnitOfMeasurement: "m³/h",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_MED"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " medium flow rate",
					EntityType:        "sensor",
					UnitOfMeasurement: "m³/h",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_HIGH"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " high flow rate",
					EntityType:        "sensor",
					UnitOfMeasurement: "m³/h",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_BOOST"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " boost flow rate",
					EntityType:        "sensor",
					UnitOfMeasurement: "m³/h",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_ECONOMY"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " economy flow rate",
					EntityType:        "sensor",
					UnitOfMeasurement: "m³/h",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_COMFORT"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " comfort flow rate",
					EntityType:        "sensor",
					UnitOfMeasurement: "m³/h",
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
				}
			}
			if write {
				m["D"+fmt.Sprint(i)+"_SPEED_LOW"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " low flow rate",
					EntityType:        "number",
					UnitOfMeasurement: "m³/h",
					Max:               250,
					Min:               100,
					Step:              10,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
					CommandTemplate:   "{{ (value / 10) | int }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_MED"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " medium flow rate",
					EntityType:        "number",
					UnitOfMeasurement: "m³/h",
					Max:               250,
					Min:               100,
					Step:              10,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
					CommandTemplate:   "{{ (value / 10) | int }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_HIGH"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " high flow rate",
					EntityType:        "number",
					UnitOfMeasurement: "m³/h",
					Max:               250,
					Min:               100,
					Step:              10,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
					CommandTemplate:   "{{ (value / 10) | int }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_BOOST"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " boost flow rate",
					EntityType:        "number",
					UnitOfMeasurement: "m³/h",
					Max:               250,
					Min:               100,
					Step:              10,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
					CommandTemplate:   "{{ (value / 10) | int }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_ECONOMY"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " economy flow rate",
					EntityType:        "number",
					UnitOfMeasurement: "m³/h",
					Max:               250,
					Min:               100,
					Step:              10,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
					CommandTemplate:   "{{ (value / 10) | int }}",
				}
				m["D"+fmt.Sprint(i)+"_SPEED_COMFORT"] = Attributes{
					Name:              "Fan " + fmt.Sprint(i) + " comfort flow rate",
					EntityType:        "number",
					UnitOfMeasurement: "m³/h",
					Max:               250,
					Min:               100,
					Step:              10,
					StateClass:        "measurement",
					ValueTemplate:     "{{ value | int * 10 }}",
					CommandTemplate:   "{{ (value / 10) | int }}",
				}
			}
		}
	}
}

func (m ParamsMap) addEnergymeters(from map[string]string, static, read, write bool) {
	for i := 1; i <= 4; i++ {
		if static {
			m["EM"+fmt.Sprint(i)+"_INSTANT"] = Attributes{
				Name:              "Energy meter " + fmt.Sprint(i) + " power",
				EntityType:        "sensor",
				DeviceClass:       "power",
				UnitOfMeasurement: "kW",
				StateClass:        "measurement",
				ValueTemplate:     "{% if value | int >= 32768 %}{{ ((value | int) - 65536 ) / 100 }}{% else %}{{ value | int / 100 }}{% endif %}",
				CommandTemplate:   "",
			}
			m["EM"+fmt.Sprint(i)+"_ACCLO"] = Attributes{
				Name:              "Energy meter " + fmt.Sprint(i) + " total energy import",
				EntityType:        "sensor",
				DeviceClass:       "energy",
				UnitOfMeasurement: "kWh",
				StateClass:        "total_increasing",
				ValueTemplate:     "{{ value | int / 10 }}",
				CommandTemplate:   "",
			}
			if i == 4 {
				m["EM"+fmt.Sprint(i)+"_ACC2LO"] = Attributes{
					Name:              "Energy meter " + fmt.Sprint(i) + " total energy export",
					EntityType:        "sensor",
					DeviceClass:       "energy",
					UnitOfMeasurement: "kWh",
					StateClass:        "total_increasing",
					ValueTemplate:     "{{ value | int / 10 }}",
					CommandTemplate:   "",
				}
			}
		}

	}
}
