package models

type Event struct {
	Device struct {
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		Name         string   `json:"name"`
		Identifiers  []string `json:"identifiers"`
	} `json:"device"`
	DeviceClass       string `json:"device_class,omitempty"`
	EntityCategory    string `json:"entity_category"`
	Name              string `json:"name"`
	StateClass        string `json:"state_class,omitempty"`
	StateTopic        string `json:"state_topic,omitempty"`
	UniqueID          string `json:"unique_id"`
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
	ValueTemplate     string `json:"value_template"`
}

func (e *Event) Init(systemID, sensorID string, attributes Attributes) {
	e.Device.Manufacturer = "Setecna"
	e.Device.Model = "REG system"
	e.Device.Name = "systemID"
	e.Device.Identifiers = []string{systemID}
	e.DeviceClass = attributes.DeviceClass
	e.EntityCategory = "diagnostic"
	e.Name = attributes.Name
	e.StateClass = attributes.StateClass
	e.StateTopic = "homeassistant/" + attributes.EntityType + "/" + systemID + "_" + sensorID
	e.UniqueID = systemID + "_" + sensorID
	e.UnitOfMeasurement = attributes.UnitOfMeasurement
	e.ValueTemplate = attributes.ValueTemplate
}
