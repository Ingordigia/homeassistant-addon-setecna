package models

type BinarySensor struct {
	Device struct {
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		Name         string   `json:"name"`
		Identifiers  []string `json:"identifiers"`
	} `json:"device"`
	DeviceClass    string `json:"device_class,omitempty"`
	EntityCategory string `json:"entity_category"`
	Name           string `json:"name"`
	PayloadOff     string `json:"payload_off"`
	PayloadOn      string `json:"payload_on"`
	StateTopic     string `json:"state_topic"`
	UniqueID       string `json:"unique_id"`
	ValueTemplate  string `json:"value_template"`
}

func (bs *BinarySensor) Init(systemID, sensorID string, attributes Attributes) {
	bs.Device.Manufacturer = "Setecna"
	bs.Device.Model = "REG system"
	bs.Device.Name = systemID
	bs.Device.Identifiers = []string{systemID}
	bs.DeviceClass = attributes.DeviceClass
	bs.EntityCategory = "diagnostic"
	bs.Name = attributes.Name
	bs.PayloadOff = "off"
	bs.PayloadOn = "on"
	bs.StateTopic = "homeassistant/" + attributes.EntityType + "/" + systemID + "_" + sensorID
	bs.UniqueID = systemID + "_" + sensorID
	bs.ValueTemplate = attributes.ValueTemplate
}
