package mqtt

import (
	"fmt"
	"log"
	"os"

	paho "github.com/eclipse/paho.mqtt.golang"
)

// NewClient creates and connects a new MQTT client using EMQX broker settings from env.
// MQTT_BROKER defaults to tcp://localhost:1883.
// MQTT_USERNAME / MQTT_PASSWORD are optional — if not set, connects anonymously.
func NewClient() paho.Client {
	broker := os.Getenv("MQTT_BROKER")
	if broker == "" {
		broker = "tcp://localhost:1883"
	}

	// Use hostname (pod name in K8s) to ensure each replica gets a unique client ID.
	// Falls back to "chatapp-backend" if hostname is unavailable.
	hostname, err := os.Hostname()
	if err != nil || hostname == "" {
		hostname = "chatapp-backend"
	}
	clientID := fmt.Sprintf("chatapp-%s", hostname)

	opts := paho.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID).
		SetCleanSession(true).
		SetOnConnectHandler(func(_ paho.Client) {
			log.Printf("[MQTT] connected to broker %s", broker)
		}).
		SetConnectionLostHandler(func(_ paho.Client, err error) {
			log.Printf("[MQTT] connection lost: %v", err)
		})

	// Only set credentials if provided in env — Mosquitto/anonymous brokers don't need them.
	if user := os.Getenv("MQTT_USERNAME"); user != "" {
		opts.SetUsername(user)
		opts.SetPassword(os.Getenv("MQTT_PASSWORD"))
	}

	client := paho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("[MQTT] WARNING: could not connect to broker %s: %v — running without MQTT", broker, token.Error())
		return client
	}

	// Subscribe to all chat topics for backend logging.
	token := client.Subscribe("chat/#", 1, func(_ paho.Client, msg paho.Message) {
		log.Printf("[MQTT LOG] Topic: %s", msg.Topic())
		log.Printf("[MQTT LOG] Payload: %s", string(msg.Payload()))
	})
	token.Wait()
	if token.Error() != nil {
		log.Printf("[MQTT] subscribe error: %v", token.Error())
	}

	return client
}
