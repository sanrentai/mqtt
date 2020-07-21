package mqtt

import (
	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

// 获取MQTT连接配置项
func GetClientOptions(conf *Config) *gomqtt.ClientOptions {
	opts := gomqtt.NewClientOptions()
	opts.SetAutoReconnect(true)
	opts.AddBroker(conf.Broker)
	if len(conf.ClientID) > 0 {
		opts.SetClientID(conf.ClientID)
	}
	return opts
}
