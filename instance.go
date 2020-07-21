package mqtt

import (
	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT连接
type Client struct {
	client      gomqtt.Client   // 实际连接
	subscribers []SubscribeType // 订阅监听器
}

// 初始化监听器
func (m *Client) Init(opts *gomqtt.ClientOptions) (err error) {
	opts.SetOnConnectHandler(m.onConnectHandler(opts.OnConnect))
	opts.SetConnectionLostHandler(m.onConnectionLostHandler(opts.OnConnectionLost))
	m.client, err = GetClient(opts)
	return
}
