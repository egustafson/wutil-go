package wlib

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-viper/mapstructure/v2"
)

type MQTTConfig struct {
	URL      string     `mapstructure:"url"`
	ClientID string     `mapstructure:"client-id,omitempty"`
	Username string     `mapstructure:"username,omitempty"`
	Password string     `mapstructure:"password,omitempty"`
	TLS      TlsProfile `mapstructure:"tls,omitempty"`
}

type MQTTOption func(*mqtt.ClientOptions)

func MQTTFactory(config *MQTTConfig, ops ...MQTTOption) mqtt.Client {
	clientOps := mqtt.NewClientOptions()
	for _, op := range ops { // apply options first
		op(clientOps)
	}

	clientOps.AddBroker(config.URL)
	if len(config.ClientID) > 0 {
		clientOps.SetClientID(config.ClientID)
	}
	if len(config.Username) > 0 {
		clientOps.SetUsername(config.Username)
	}
	if len(config.Password) > 0 {
		clientOps.SetPassword(config.Password)
	}
	if tlsConfig := config.TLS.TlsConfig(); tlsConfig != nil {
		clientOps.SetTLSConfig(tlsConfig)
	}

	return mqtt.NewClient(clientOps)
}

func WithClientOptions(o *mqtt.ClientOptions) MQTTOption {
	return func(opt *mqtt.ClientOptions) {
		*opt = *o // copy struct fields
	}
}
