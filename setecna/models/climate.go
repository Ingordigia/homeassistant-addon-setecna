package models

type Climate struct {
	Name     string   `json:"name"`
	UniqueID string   `json:"unique_id"`
	Modes    []string `json:"modes"`
	Device   struct {
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		Name         string   `json:"name"`
		Identifiers  []string `json:"identifiers"`
	} `json:"device"`
	EntityCategory              string  `json:"entity_category"`
	ModeStateTopic              string  `json:"mode_state_topic"`
	ModeStateTemplate           string  `json:"mode_state_template"`
	CurrentHumidityTopic        string  `json:"current_humidity_topic"`
	CurrentHumidityTemplate     string  `json:"current_humidity_template"`
	TargetHumidityStateTopic    string  `json:"target_humidity_state_topic"`
	TargetHumidityCommandTopic  string  `json:"target_humidity_command_topic"`
	MinHumidity                 int     `json:"min_humidity"`
	MaxHumidity                 int     `json:"max_humidity"`
	CurrentTemperatureTopic     string  `json:"current_temperature_topic"`
	CurrentTemperatureTemplate  string  `json:"current_temperature_template"`
	TemperatureLowStateTopic    string  `json:"temperature_low_state_topic"`
	TemperatureLowCommandTopic  string  `json:"temperature_low_command_topic"`
	TemperatureHighStateTopic   string  `json:"temperature_high_state_topic"`
	TemperatureHighCommandTopic string  `json:"temperature_high_command_topic"`
	TempStep                    float64 `json:"temp_step"`
	MinTemp                     int     `json:"min_temp"`
	MaxTemp                     int     `json:"max_temp"`
}
