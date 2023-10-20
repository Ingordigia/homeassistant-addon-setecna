package mqtt

import (
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type Message struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
	Qos     int    `json:"qos"`
}

type Messages []Message

type MqttServer struct {
	client MQTT.Client
}

func (s *MqttServer) Connect(host, user, password string) {
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://" + host + ":1883")
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetClientID("SetecnaScraperAddon")

	s.client = MQTT.NewClient(opts)
	if token := s.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func (s *MqttServer) Disconnect() {
	s.client.Disconnect(250)
}

func (s *MqttServer) BatchPublish(u Messages, delay int64) {
	for _, m := range u {
		// fmt.Println("---- Publishing message \"" + m.Message + "\" on topic \"" + m.Topic + "\"")
		token := s.client.Publish(m.Topic, byte(m.Qos), false, m.Message)
		token.Wait()
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}

func (s *MqttServer) Publish(m Message) {
	// fmt.Println("---- Publishing message \"" + m.Message + "\" on topic \"" + m.Topic + "\"")
	token := s.client.Publish(m.Topic, byte(m.Qos), false, m.Message)
	token.Wait()
}

func (s *MqttServer) ListenForChange(topic string, handlerFunction MQTT.MessageHandler) {
	s.client.Subscribe(topic, 0, handlerFunction)
}
