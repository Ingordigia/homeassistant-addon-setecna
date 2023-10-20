package models

// https://github.com/home-assistant/core/issues/84745
// https://github.com/home-assistant/core/issues/85187
// https://github.com/home-assistant/frontend/pull/14502
// https://github.com/home-assistant/architecture/discussions/856

import (
	"fmt"

	"github.com/Ingordigia/homeassistant-addon-setecna/pkg/helpers"
)

type ClimateWithHumidity struct {
	// ActionTemplate string `json:"action_template"`
	// ActionTopic    string `json:"action_topic"`
	Name     string `json:"name"`
	UniqueID string `json:"unique_id"`
	Device   struct {
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		Name         string   `json:"name"`
		Identifiers  []string `json:"identifiers"`
	} `json:"device"`
	EntityCategory                 string   `json:"entity_category"`
	ModeStateTopic                 string   `json:"mode_state_topic"`
	ModeStateTemplate              string   `json:"mode_state_template"`
	Modes                          []string `json:"modes"`
	CurrentHumidityTopic           string   `json:"current_humidity_topic"`
	CurrentHumidityTemplate        string   `json:"current_humidity_template"`
	TargetHumidityStateTopic       string   `json:"target_humidity_state_topic"`
	TargetHumidityStateTemplate    string   `json:"target_humidity_state_template"`
	TargetHumidityCommandTopic     string   `json:"target_humidity_command_topic"`
	TargetHumidityCommandTemplate  string   `json:"target_humidity_command_template"`
	MinHumidity                    int      `json:"min_humidity"`
	MaxHumidity                    int      `json:"max_humidity"`
	CurrentTemperatureTopic        string   `json:"current_temperature_topic"`
	CurrentTemperatureTemplate     string   `json:"current_temperature_template"`
	PresetModeCommandTemplate      string   `json:"preset_mode_command_template"`
	PresetModeCommandTopic         string   `json:"preset_mode_command_topic"`
	PresetModeStateTopic           string   `json:"preset_mode_state_topic"`
	PresetModeValueTemplate        string   `json:"preset_mode_value_template"`
	PresetModes                    []string `json:"preset_modes"`
	TemperatureLowStateTopic       string   `json:"temperature_low_state_topic"`
	TemperatureLowStateTemplate    string   `json:"temperature_low_state_template"`
	TemperatureLowCommandTopic     string   `json:"temperature_low_command_topic"`
	TemperatureLowCommandTemplate  string   `json:"temperature_low_command_template"`
	TemperatureHighStateTopic      string   `json:"temperature_high_state_topic"`
	TemperatureHighStateTemplate   string   `json:"temperature_high_state_template"`
	TemperatureHighCommandTopic    string   `json:"temperature_high_command_topic"`
	TemperatureHighCommandTemplate string   `json:"temperature_high_command_template"`
	TempStep                       float64  `json:"temp_step"`
	MinTemp                        int      `json:"min_temp"`
	MaxTemp                        int      `json:"max_temp"`
}

func (c *ClimateWithHumidity) Init(number int, systemID string, season helpers.Season) {
	// c.ActionTemplate = "homeassistant/binary_sensor/" + systemID + "_Z" + fmt.Sprint(number) + "_OUTPUT"
	// if season == helpers.Summer {
	// 	c.ActionTopic = "{% if value == \"1\" %}cooling{% else %}off{% endif %}"
	// } else {
	// 	c.ActionTopic = "{% if value == \"1\" %}heating{% else %}off{% endif %}"
	// }
	c.Name = "Zone " + fmt.Sprint(number)
	c.UniqueID = systemID + "_zone_" + fmt.Sprint(number)
	c.Device.Manufacturer = "Setecna"
	c.Device.Model = "REG system"
	c.Device.Name = systemID
	c.Device.Identifiers = []string{systemID}
	c.EntityCategory = "config"
	c.ModeStateTopic = "homeassistant/binary_sensor/" + systemID + "_Z" + fmt.Sprint(number) + "_OUTPUT"
	if season == helpers.Summer {
		c.ModeStateTemplate = "{% if value == \"1\" %}cool{% else %}off{% endif %}"
		c.Modes = []string{"cool", "off"}
		c.TemperatureHighStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_ES"
		c.TemperatureHighStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureHighCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_ES/set"
		c.TemperatureHighCommandTemplate = "{{ value | int * 10 }}"
		c.TemperatureLowStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CS"
		c.TemperatureLowStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureLowCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CS/set"
		c.TemperatureLowCommandTemplate = "{{ value | int * 10 }}"
	} else {
		c.ModeStateTemplate = "{% if value == \"1\" %}heat{% else %}off{% endif %}"
		c.Modes = []string{"heat", "off"}
		c.TemperatureLowStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_EW"
		c.TemperatureLowStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureLowCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_EW/set"
		c.TemperatureLowCommandTemplate = "{{ value | int * 10 }}"
		c.TemperatureHighStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CW"
		c.TemperatureHighStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureHighCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CW/set"
		c.TemperatureHighCommandTemplate = "{{ value | int * 10 }}"
	}
	// c.ModeStateTopic = "homeassistant/sensor/" + systemID + "_GLOBAL_SEASON"
	// c.ModeStateTemplate = "{% if value == \"1\" %}cool{% else %}off{% endif %}"
	// c.Modes = []string{}
	c.CurrentHumidityTopic = "homeassistant/sensor/" + systemID + "_Z" + fmt.Sprint(number) + "_RH"
	c.CurrentHumidityTemplate = "{{ value | int / 10 }}"
	c.CurrentTemperatureTopic = "homeassistant/sensor/" + systemID + "_Z" + fmt.Sprint(number) + "_TEMP"
	c.CurrentTemperatureTemplate = "{{ value | int / 10 }}"
	c.MinHumidity = 45
	c.MinTemp = 15
	c.MaxHumidity = 75
	c.MaxTemp = 30
	c.PresetModeCommandTopic = "homeassistant/select/" + systemID + "_Z" + fmt.Sprint(number) + "_FORCING/set"
	c.PresetModeCommandTemplate = "{% if value == \"forced off\" %}1{% elif value == \"forced economy\" %}2{% elif value == \"forced comfort\" %}3{% else %}0{% endif %}"
	c.PresetModeStateTopic = "homeassistant/select/" + systemID + "_Z" + fmt.Sprint(number) + "_FORCING"
	c.PresetModeValueTemplate = "{% if value == \"1\" %}forced off{% elif value == \"2\" %}forced economy{% elif value == \"3\" %}forced comfort{% else %}none{% endif %}"
	c.PresetModes = []string{"forced off", "forced economy", "forced comfort"}
	c.TargetHumidityStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_RH"
	c.TargetHumidityStateTemplate = "{{ value | int / 10 }}"
	c.TargetHumidityCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_RH/set"
	c.TargetHumidityCommandTemplate = "{{ value | int * 10 }}"
	c.TempStep = 0.5
}

type ClimateWithoutHumidity struct {
	// ActionTemplate string `json:"action_template"`
	// ActionTopic    string `json:"action_topic"`
	Name     string `json:"name"`
	UniqueID string `json:"unique_id"`
	Device   struct {
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		Name         string   `json:"name"`
		Identifiers  []string `json:"identifiers"`
	} `json:"device"`
	EntityCategory                 string   `json:"entity_category"`
	ModeStateTopic                 string   `json:"mode_state_topic"`
	ModeStateTemplate              string   `json:"mode_state_template"`
	Modes                          []string `json:"modes"`
	CurrentTemperatureTopic        string   `json:"current_temperature_topic"`
	CurrentTemperatureTemplate     string   `json:"current_temperature_template"`
	PresetModeCommandTemplate      string   `json:"preset_mode_command_template"`
	PresetModeCommandTopic         string   `json:"preset_mode_command_topic"`
	PresetModeStateTopic           string   `json:"preset_mode_state_topic"`
	PresetModeValueTemplate        string   `json:"preset_mode_value_template"`
	PresetModes                    []string `json:"preset_modes"`
	TemperatureLowStateTopic       string   `json:"temperature_low_state_topic"`
	TemperatureLowStateTemplate    string   `json:"temperature_low_state_template"`
	TemperatureLowCommandTopic     string   `json:"temperature_low_command_topic"`
	TemperatureLowCommandTemplate  string   `json:"temperature_low_command_template"`
	TemperatureHighStateTopic      string   `json:"temperature_high_state_topic"`
	TemperatureHighStateTemplate   string   `json:"temperature_high_state_template"`
	TemperatureHighCommandTopic    string   `json:"temperature_high_command_topic"`
	TemperatureHighCommandTemplate string   `json:"temperature_high_command_template"`
	TempStep                       float64  `json:"temp_step"`
	MinTemp                        int      `json:"min_temp"`
	MaxTemp                        int      `json:"max_temp"`
}

func (c *ClimateWithoutHumidity) Init(number int, systemID string, season helpers.Season) {
	// c.ActionTemplate = "homeassistant/binary_sensor/" + systemID + "_Z" + fmt.Sprint(number) + "_OUTPUT"
	// if season == helpers.Summer {
	// 	c.ActionTopic = "{% if value == \"1\" %}cooling{% else %}off{% endif %}"
	// } else {
	// 	c.ActionTopic = "{% if value == \"1\" %}heating{% else %}off{% endif %}"
	// }
	c.Name = "Zone " + fmt.Sprint(number)
	c.UniqueID = "zone_" + fmt.Sprint(number)
	c.Device.Manufacturer = "Setecna"
	c.Device.Model = "REG system"
	c.Device.Name = systemID
	c.Device.Identifiers = []string{systemID}
	c.EntityCategory = "config"
	c.ModeStateTopic = "homeassistant/binary_sensor/" + systemID + "_Z" + fmt.Sprint(number) + "_OUTPUT"
	if season == helpers.Summer {
		c.ModeStateTemplate = "{% if value == \"1\" %}cool{% else %}off{% endif %}"
		c.Modes = []string{"cool", "off"}
		c.TemperatureLowStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CS"
		c.TemperatureLowStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureLowCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CS/set"
		c.TemperatureLowCommandTemplate = "{{ value | int * 10 }}"
		c.TemperatureHighStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_ES"
		c.TemperatureHighStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureHighCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_ES/set"
		c.TemperatureHighCommandTemplate = "{{ value | int * 10 }}"
	} else {
		c.ModeStateTemplate = "{% if value == \"1\" %}heat{% else %}off{% endif %}"
		c.Modes = []string{"heat", "off"}
		c.TemperatureHighStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CW"
		c.TemperatureHighStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureHighCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_CW/set"
		c.TemperatureHighCommandTemplate = "{{ value | int * 10 }}"
		c.TemperatureLowStateTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_EW"
		c.TemperatureLowStateTemplate = "{{ value | int / 10 }}"
		c.TemperatureLowCommandTopic = "homeassistant/number/" + systemID + "_Z" + fmt.Sprint(number) + "_SET_EW/set"
		c.TemperatureLowCommandTemplate = "{{ value | int * 10 }}"
	}
	// c.Modes = []string{"cool", "heat", "off"}
	// // c.ModeStateTopic = "homeassistant/sensor/" + systemID + "_GLOBAL_SEASON"
	// // c.ModeStateTemplate = "{% if value == \"1\" %}cool{% else %}heat{% endif %}"
	// c.Modes = []string{}
	c.CurrentTemperatureTopic = "homeassistant/sensor/" + systemID + "_Z" + fmt.Sprint(number) + "_TEMP"
	c.CurrentTemperatureTemplate = "{{ value | int / 10 }}"
	c.PresetModeCommandTopic = "homeassistant/select/" + systemID + "_Z" + fmt.Sprint(number) + "_FORCING/set"
	c.PresetModeCommandTemplate = "{% if value == \"forced off\" %}1{% elif value == \"forced economy\" %}2{% elif value == \"forced comfort\" %}3{% else %}0{% endif %}"
	c.PresetModeStateTopic = "homeassistant/select/" + systemID + "_Z" + fmt.Sprint(number) + "_FORCING"
	c.PresetModeValueTemplate = "{% if value == \"1\" %}forced off{% elif value == \"2\" %}forced economy{% elif value == \"3\" %}forced comfort{% else %}none{% endif %}"
	c.PresetModes = []string{"forced off", "forced economy", "forced comfort"}
	c.TempStep = 0.5
	c.MinTemp = 15
	c.MaxTemp = 30
}
