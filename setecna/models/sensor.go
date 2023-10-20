package models

type Sensor struct {
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
	ValueTemplate     string `json:"value_template,omitempty"`
}

func (s *Sensor) Init(systemID, sensorID string, attributes Attributes) {
	s.Device.Manufacturer = "Setecna"
	s.Device.Model = "REG system"
	s.Device.Name = systemID
	s.Device.Identifiers = []string{systemID}
	s.DeviceClass = attributes.DeviceClass
	s.EntityCategory = "diagnostic"
	s.Name = attributes.Name
	s.StateClass = attributes.StateClass
	s.StateTopic = "homeassistant/" + attributes.EntityType + "/" + systemID + "_" + sensorID
	s.UniqueID = systemID + "_" + sensorID
	s.UnitOfMeasurement = attributes.UnitOfMeasurement
	s.ValueTemplate = attributes.ValueTemplate
}
