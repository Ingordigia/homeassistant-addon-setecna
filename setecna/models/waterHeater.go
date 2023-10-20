package models

// TODO: valutare se usare un climate per supportare 2 temperature target
type WaterHeater struct {
	CurrentTemperatureTemplate string `json:" current_temperature_template "`
	CurrentTemperatureTopic    string `json:"current_temperature_topic"`
	Device                     struct {
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		Name         string   `json:"name"`
		Identifiers  []string `json:"identifiers"`
	} `json:"device"`
	EntityCategory             string   `json:"entity_category"`
	MaxTemp                    int      `json:"max_temp"`
	MinTemp                    int      `json:"min_temp"`
	ModeCommandTemplate        string   `json:"mode_command_template"`
	ModeCommandTopic           string   `json:"mode_command_topic"`
	ModeStateTemplate          string   `json:"mode_state_template"`
	ModeStateTopic             string   `json:"mode_state_topic"`
	Modes                      []string `json:"modes"`
	Name                       string   `json:"name"`
	PayloadOff                 string   `json:"payload_off"`
	PayloadOn                  string   `json:"payload_on"`
	PowerCommandTemplate       string   `json:"power_command_template"`
	PowerCommandTopic          string   `json:"power_command_topic"`
	Precision                  float64  `json:"precision"`
	TemperatureCommandTemplate string   `json:"temperature_command_template"`
	TemperatureCommandTopic    string   `json:"temperature_command_topic"`
	TemperatureStateTemplate   string   `json:"temperature_state_template"`
	TemperatureStateTopic      string   `json:"temperature_state_topic"`
	TemperatureUnit            string   `json:"temperature_unit"`
	UniqueID                   string   `json:"unique_id"`
	ValueTemplate              string   `json:"value_template"`
}

func (wh *WaterHeater) Init(systemID string) {
	wh.CurrentTemperatureTemplate = "{{ value | int / 10 }}"
	wh.CurrentTemperatureTopic = "setecna/GLOBAL_T_ACS"
	wh.Device.Manufacturer = "Setecna"
	wh.Device.Model = "REG system"
	wh.Device.Name = systemID
	wh.Device.Identifiers = []string{systemID}
	wh.EntityCategory = "config"
	wh.MaxTemp = 60
	wh.MinTemp = 30
	wh.ModeCommandTemplate = "{% if value == \"off\" %}1{% elif value == \"eco\" %}2{% elif value == \"high_demand\" %}3{% elif value == \"heat_pump\" %}0{% endif %}"
	wh.ModeCommandTopic = "setecna/MT8_FORCING"
	wh.ModeStateTemplate = "{% if value == \"1\" %}off{% elif value == \"2\" %}eco{% elif value == \"3\" %}high_demand{% endif %}"
	wh.ModeStateTopic = "setecna/MT8_MODE"
	wh.Modes = []string{"off", "eco", "high_demand", "heat_pump"}
	wh.Name = "WaterHeater"
	// wh.PayloadOff = ""
	// wh.PayloadOn = ""
	// wh.PowerCommandTemplate = ""
	// wh.PowerCommandTopic = ""
	wh.Precision = 0.5
	// wh.TemperatureCommandTemplate = ""
	// wh.TemperatureCommandTopic = ""
	// wh.TemperatureStateTemplate = ""
	// wh.TemperatureStateTopic = ""
	wh.TemperatureUnit = "°C"
	wh.UniqueID = "WaterHeater"
	wh.ValueTemplate = ""

}
