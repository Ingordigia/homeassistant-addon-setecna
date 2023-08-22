package models

type Response struct {
	Status    string `json:"Status"`
	Timestamp string `json:"Timestamp"`
	Latest    string `json:"Latest"`
	Rows      int    `json:"Rows"`
	Data      []struct {
		ID string `json:"Id"`
		V  int    `json:"V"`
	} `json:"Data"`
}

// type entity struct {
// 	Id          string  `json:"id"`
// 	Name        string  `json:"name"`
// 	Type        string  `json:"type"`
// 	MqttTopic   string  `json:"mqtt_topic"`
// 	FloatValue  float64 `json:"int_value"`
// 	StringValue string  `json:"string_value"`
// }

// type System struct {
// 	system map[string]entity
// }
