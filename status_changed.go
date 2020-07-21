package mqtt

import (
	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

// 连接上服务器的操作
func (m *Client) onConnectHandler(handler gomqtt.OnConnectHandler) gomqtt.OnConnectHandler {
	return func(c gomqtt.Client) {
		for _, item := range m.subscribers {
			m.subscribe(item)
		}
		handler(c)
	}
}

// 丢失连接的操作(自动重连)
func (m *Client) onConnectionLostHandler(handler gomqtt.ConnectionLostHandler) gomqtt.ConnectionLostHandler {
	return func(c gomqtt.Client, e error) {
		handler(c, e)
	}
}
