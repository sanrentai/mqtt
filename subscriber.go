package mqtt

import (
	"time"

	gomqtt "github.com/eclipse/paho.mqtt.golang"
)

type SubscribeType struct {
	Topic      string
	Qos        byte
	Callback   gomqtt.MessageHandler
	RetryTimes int // 为0表示无限重试
}

// 注册订阅消息
func (m *Client) Subscribe(item SubscribeType) {
	m.subscribers = append(m.subscribers, item)
}

func (m *Client) subscribe(item SubscribeType) {
	times := 0
	for {
		token, err := m.subscribeItem(item)
		if err != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			} else {
				panic(err)
			}
		}
		if token.Wait() && token.Error() != nil {
			if item.RetryTimes == 0 || times < item.RetryTimes {
				times++
				time.Sleep(3 * time.Second)
				continue
			} else {
				panic(token.Error())
			}
		}
		break
	}
}

func (m Client) subscribeItem(item SubscribeType) (token gomqtt.Token, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
		return
	}()
	token = m.client.Subscribe(item.Topic, item.Qos, item.Callback)
	return
}
